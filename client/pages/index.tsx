import { useState, useEffect, useContext } from "react";
import { API_URL, WEBSOCKET_URL } from "../constants";
import { v4 as uuidv4 } from "uuid";
import { AuthContext } from "../modules/auth_provider";
import { WebsocketContext } from "../modules/websocket_provider";
import { useRouter } from "next/router";

const index = () => {
  const [rooms, setRooms] = useState<{ id: string; name: string }[]>([]);
  const [roomName, setRoomName] = useState("");
  const { user } = useContext(AuthContext);
  const { setConn } = useContext(WebsocketContext);

  const router = useRouter();

  const getRooms = async () => {
    try {
      const res = await fetch(`http://localhost:8081/api/chats/getRooms`, {
        method: "GET",
      });

      const data = await res.json();
      if (res.ok) {
        setRooms(data);
      }
    } catch (err) {
      console.log(err);
    }
  };

  useEffect(() => {
    getRooms();
  }, []);

  const submitHandler = async (e: React.SyntheticEvent) => {
    e.preventDefault();

    try {
      const res = await fetch(`http://localhost:8081/api/chats/createRoom`, {
        method: "POST",
        // credentials: "include",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({
          id: uuidv4(),
          name: roomName,
        }),
      });

      if (res.ok) {
        getRooms();
      }
    } catch (err) {
      console.log(err);
    }
  };

  const joinRoom = async (roomId: string) => {
    const ws = new WebSocket(
      `ws://localhost:8081/api/chats/joinRoom/${roomId}?userId=${user.userId}&username=${user.username}`
    );
    if (ws.OPEN) {
      setConn(ws);
      router.push("/app");
      return;
    }
  };

  return (
    <>
      <div className="w-full h-full px-4 my-8 md:mx-32">
        <div className="flex justify-center p-5 mt-3">
          <input
            type="text"
            className="p-2 border rounded-md border-grey focus:outline-none focus:border-blue"
            placeholder="room name"
            value={roomName}
            onChange={(e) => setRoomName(e.target.value)}
          />
          <button
            className="p-2 text-white border rounded-md bg-blue md:ml-4"
            onClick={submitHandler}
          >
            create room
          </button>
        </div>
        <div className="mt-6">
          <div className="font-bold">Available Rooms</div>
          <div className="grid grid-cols-1 gap-4 mt-6 md:grid-cols-5">
            {rooms.map((room, index) => (
              <div
                key={index}
                className="flex items-center w-full p-4 border rounded-md border-blue"
              >
                <div className="w-full">
                  <div className="text-sm">room</div>
                  <div className="text-lg font-bold text-blue">{room.name}</div>
                </div>
                <div className="">
                  <button
                    className="px-4 text-white rounded-md bg-blue"
                    onClick={() => joinRoom(room.id)}
                  >
                    join
                  </button>
                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </>
  );
};

export default index;
