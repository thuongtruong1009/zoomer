const config = {
    iceServers: [{ urls: "stun:stun.l.google.com:19302" }],
};

const connections = new Map(); // Map of connections for each peer
const videos = new Map(); // Map of video elements for each peer
const localVideo = document.createElement("video"); // Local video element
localVideo.autoplay = true;
localVideo.muted = true;
document.body.appendChild(localVideo);

// Connect to RabbitMQ
const amqp = require("amqplib/callback_api");
const mqUrl = "amqp://guest:guest@localhost:5672";
amqp.connect(mqUrl, (err, conn) => {
            if (err) {
                throw err;
            }
            conn.createChannel((err, ch) => {
                        if (err) {
                            throw err;
                        }

                        // Declare the exchange for sending messages
                        const ex = "webrtc-exchange";
                        ch.assertExchange(ex, "fanout", { durable: false });

                        // Declare a queue for each peer
                        const q = "webrtc-queue";
                        ch.assertQueue(q, { exclusive: true }, (err, q) => {
                                    if (err) {
                                        throw err;
                                    }
                                    ch.bindQueue(q.queue, ex, "");

                                    // Add tracks to the peer connection from the local video element
                                    navigator.mediaDevices.getUserMedia({ video: true, audio: true }).then((stream) => {
                                        stream.getTracks().forEach((track) => {
                                            connections.forEach((conn) => {
                                                conn.addTrack(track, stream);
                                            });
                                        });
                                        localVideo.srcObject = stream;

                                        // Send an offer to each peer
                                        connections.forEach((conn, peerId) => {
                                            conn.createOffer().then((offer) => {
                                                conn.setLocalDescription(offer).then(() => {
                                                    const message = {
                                                        type: "offer",
                                                        sdp: conn.localDescription.sdp,
                                                        from: "local",
                                                        to: peerId,
                                                    };
                                                    ch.publish(ex, "", Buffer.from(JSON.stringify(message)));
                                                });
                                            });
                                        });
                                    });

                                    // Receive messages from the queue
                                    ch.consume(q.queue, (msg) => {
                                                const message = JSON.parse(msg.content.toString());
                                                const peerId = message.from;
                                                if (!connections.has(peerId)) {
                                                    // Create a new connection for the peer
                                                    const connection = new RTCPeerConnection(config);
                                                    connection.onicecandidate = (event) => {
                                                        if (event.candidate) {
                                                            const message = {
                                                                type: "icecandidate",
                                                                candidate: event.candidate.candidate,
                                                                sdpMid: event.candidate.sdpMid,
                                                                sdpMLineIndex: event.candidate.sdpMLineIndex,
                                                                from: "local",
                                                                to: peerId,
                                                            };
                                                            ch.publish(ex, "", Buffer.from(JSON.stringify(message)));
                                                        }
                                                    };
                                                    connection.ontrack = (event) => {
                                                        if (event.track.kind === "video") {
                                                            videos.get(peerId).srcObject = event.streams[0];
                                                        }
                                                    };
                                                    connections.set(peerId, connection);

                                                    // Create a new video element for the peer
                                                    const video = document.createElement("video");
                                                    video.autoplay = true;
                                                    document.body.appendChild(video);
                                                    videos.set(peerId, video);
                                                }
                                                switch (message.type) {
                                                    case "offer":
                                                        // Set the remote description from the