import { configureStore } from '@reduxjs/toolkit';
import myReducer from './index';
import { useSelector } from 'react-redux';

export const store = configureStore({
  reducer: {
    my: myReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
