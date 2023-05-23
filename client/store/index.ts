import { createSlice, PayloadAction } from '@reduxjs/toolkit';

type Contact = {
  username: string;
  last_activity: number
}
interface MyState {
  items: Contact[];
}

const initialState: MyState = {
  items: [],
};

const mySlice = createSlice({
  name: 'contactReducer',
  initialState,
  reducers: {
    addItem: (state, action: PayloadAction<Contact>) => {
      state.items.unshift(action.payload);
    },
    addAll: (state, action: PayloadAction<Contact[]>) => {
      state.items = action.payload;
    },
  },

});

export const { addItem } = mySlice.actions;
export const { addAll } = mySlice.actions;

export default mySlice.reducer;
