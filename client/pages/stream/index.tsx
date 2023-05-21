const WindowOptions = () => {
  const width = 1000;
  const height = 600;
  const left = (window.innerWidth / 2) - (width / 2);
  const top = (window.innerHeight / 2) - (height / 2);
  const size = 'width=' + width + ', height=' + height + ', left=' + left + ', top=' + top;
  const options = size + ', toolbar=yes, scrollbars=yes, resizable=yes';
  return options;
}

const CreateRoom: React.FC = () => {
  const create = async (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault();

    const resp = await fetch("http://localhost:8081/create");
    const { data } = await resp.json();

    window.open(`/stream/${data.room_id}`, '_blank', WindowOptions())
  };

  return (
    <div>
      <button onClick={create}>Create Room</button>
    </div>
  );
};

export default CreateRoom;
