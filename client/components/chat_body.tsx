import React from "react";
import { Message } from "../pages/app";

const ChatBody = ({ data }: { data: Array<Message> }) => {
  return (
    <>
      {data.map((message: Message, index: number) => {
        if (message.type == "self") {
          return (
            <div
              className="flex flex-col justify-end w-full mt-2 text-right"
              key={index}
            >
              <div className="text-sm">{message.username}</div>
              <div>
                <div className="inline-block px-4 py-1 mt-1 text-white rounded-md bg-blue">
                  {message.content}
                </div>
              </div>
            </div>
          );
        } else if (message.type == "recv") {
          return (
            <div className="mt-2" key={index}>
              <div className="text-sm">{message.username}</div>
              <div>
                <div className="inline-block px-4 py-1 mt-1 rounded-md bg-grey text-dark-secondary">
                  {message.content}
                </div>
              </div>
            </div>
          );
        } else {
          return (
            <div className="text-center" key={index}>
              <div className="inline-block my-1 text-gray-400">
                {message.content}
              </div>
            </div>
          );
        }
      })}
    </>
  );
};

export default ChatBody;
