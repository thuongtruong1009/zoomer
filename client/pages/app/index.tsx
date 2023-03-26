import React, { useState, useRef, useContext, useEffect } from "react";
import ChatBody from "../../components/chat_body";
import { WebsocketContext } from "../../modules/websocket_provider";
import { useRouter } from "next/router";
import { API_URL } from "../../constants";
import autosize from "autosize";
import { AuthContext } from "../../modules/auth_provider";

export type Message = {
  content: string;
  client_id: string;
  username: string;
  room_id: string;
  type: "recv" | "self" | "other";
};

const index = () => {
  const [messages, setMessage] = useState<Array<Message>>([]);
  const textarea = useRef<HTMLTextAreaElement>(null);
  const { conn } = useContext(WebsocketContext);
  const [users, setUsers] = useState<Array<{ username: string }>>([]);
  const { user } = useContext(AuthContext);

  const router = useRouter();

  useEffect(() => {
    if (conn === null) {
      router.push("/");
      return;
    }

    const roomId = conn.url.split("/")[6].split("?")[0];

    async function getUsers() {
      try {
        const res = await fetch(
          `http://localhost:8081/api/chats/getClients/${roomId}`,
          {
            method: "GET",
            headers: { "Content-Type": "application/json" },
          }
        );
        const data = await res.json();
        console.log("res", data);

        setUsers(data);
      } catch (e) {
        console.error(e);
      }
    }
    getUsers();
  }, []);

  useEffect(() => {
    if (textarea.current) {
      autosize(textarea.current);
    }

    if (conn === null) {
      router.push("/");
      return;
    }

    conn.onmessage = (message) => {
      const m: Message = JSON.parse(message.data);
      if (m.content.includes("joined")) {
        setUsers([...users, { username: m.username }]);
      }

      if (m.content.includes("left")) {
        const deleteUser = users.filter((user) => user.username != m.username);
        setUsers([...deleteUser]);
        setMessage([...messages, m]);
        return;
      }

      let msgType: string = m.type;

      // user?.username == m.username ? (m.type = "self") : (m.type = "recv");
      if (msgType == "text") {
        user?.username == m.username ? (m.type = "self") : (m.type = "recv");
        // if (user?.username == m.username) {
        //   m.type = "self";
        // } else if (user?.username != m.username) {
        //   m.type = "recv";
        // }
      } else {
        m.type = "other";
      }
      setMessage([...messages, m]);
    };

    conn.onclose = () => {};
    conn.onerror = () => {};
    conn.onopen = () => {};
  }, [textarea, messages, conn, users]);

  const sendMessage = () => {
    if (!textarea.current?.value) return;
    if (conn === null) {
      router.push("/");
      return;
    }

    conn.send(textarea.current.value);
    textarea.current.value = "";
  };

  return (
    <>
      <div className="flex flex-col w-full">
        <div className="p-4 md:mx-6 mb-14">
          <ChatBody data={messages} />
        </div>
        <div className="fixed bottom-0 w-full mt-4">
          <div className="flex px-4 py-2 rounded-md md:flex-row bg-grey md:mx-4">
            <div className="flex w-full mr-4 border rounded-md border-blue">
              <textarea
                ref={textarea}
                placeholder="type your message here"
                className="w-full h-10 p-2 rounded-md focus:outline-none"
                style={{ resize: "none" }}
              />
            </div>
            <div className="flex items-center">
              <button
                className="p-2 text-white rounded-md bg-blue"
                onClick={sendMessage}
              >
                Send
              </button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
};

export default index;
