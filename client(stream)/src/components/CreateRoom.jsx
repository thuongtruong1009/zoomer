import React from "react";
import { useNavigate } from "react-router-dom";

const CreateRoom = () => {
  const navigate = useNavigate();

  const create = async (e) => {
    e.preventDefault();

    const resp = await fetch("http://localhost:8081/create");
    const data = await resp.json();

    navigate(`/room/${data.data.room_id}`);
  };

  return (
    <div>
      <button onClick={create}>Create Room</button>
    </div>
  );
};

export default CreateRoom;
