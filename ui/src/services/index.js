import axios from 'axios';
import { storage } from '../utils';

// axios.defaults.headers.common["X-Requested-With"] = ["XMLHttpRequest"];

const instance = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    Accept: 'application/json',
    Authorization: `Bearer ${storage.getItem('token')}`,
  },
});

instance.getErrorMessageFromResponse = (error) => {
  return (
    error?.response?.data ||
    error?.data ||
    'Something went wrong when fetching data from the API'
  );
};

export default instance;
