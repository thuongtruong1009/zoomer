import { configureStore } from '@reduxjs/toolkit';
import myReducer from './index';
import { useDispatch } from 'react-redux';

export const store = configureStore({
  reducer: {
    contactReducer: myReducer,
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export const useAppDispatch = () => useDispatch<AppDispatch>();
