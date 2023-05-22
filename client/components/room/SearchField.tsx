import * as React from 'react';
import TextField from '@mui/material/TextField';
import Autocomplete, { createFilterOptions } from '@mui/material/Autocomplete';
import { Fragment, useEffect, useState } from 'react';

interface UserOptionType {
  username: string;
}

const users: readonly UserOptionType[] = [
  { username: 'John' },
  { username: 'Oliver' },
  { username: 'Harry' },
  { username: 'George' },
  { username: 'Noah' },
];

const filter = createFilterOptions<UserOptionType>();

interface Props {
  handleOk: (value: string) => void
}

export const SearchField: React.FC<Props> = ({handleOk}: Props) => {
  const [value, setValue] = useState<UserOptionType | null>(null);

  const [dialogValue, setDialogValue] = useState({
    username: '',
  });

  useEffect(() => {
    if (dialogValue && dialogValue.username) {
      handleOk(dialogValue.username)
    }
  }, [dialogValue])

  return (
    <Fragment>
      <Autocomplete
        value={value}
        onChange={(event, newValue) => {
          if (typeof newValue === 'string') {
            setTimeout(() => {
              setDialogValue({
                username: newValue,
              });
            });
          } else if (newValue && newValue.username) {
            setDialogValue({
              username: newValue.username,
            });
          } else {
            setValue(newValue);
          }
        }}
        filterOptions={(options, params) => {
          const filtered = filter(options, params);

          if (params.inputValue !== '') {
            filtered.push({
              username: params.inputValue,
            });
          }

          return filtered;
        }}
        id="free-solo-dialog-demo"
        options={users}
        getOptionLabel={(option) => {
          if (typeof option === 'string') {
            return option;
          }
          if (option.username) {
            return option.username;
          }
          return option.username;
        }}
        selectOnFocus
        clearOnBlur
        handleHomeEndKeys
        renderOption={(props, option) => <li {...props}>{option.username}</li>}
        sx={{ width: '100%', height: 350 , justifySelfContent: 'center'}}
        freeSolo
        renderInput={(params) => <TextField {...params} label="Enter your friend name" />}
      />
    </Fragment>
  );
}
