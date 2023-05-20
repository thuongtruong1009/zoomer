import { useRouter } from 'next/router';

const CreateRoom: React.FC = () => {
  const router = useRouter();

  const create = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();

    const resp = await fetch("http://localhost:8081/create");
    const { data } = await resp.json();

    router.push(`/stream/${data.room_id}`);
  };

  return (
    <div>
      <button onClick={create}>Create Room</button>
    </div>
  );
};

export default CreateRoom;
