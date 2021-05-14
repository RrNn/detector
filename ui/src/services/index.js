import axios from 'axios';

// axios.defaults.headers.common['X-Requested-With'] = ['XMLHttpRequest'];

const instance = axios.create({
  baseURL: '/api',
  timeout: 10000,
  headers: {
    Accept: 'application/json',
    Authorization: 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZGRyZXNzIjoiS2FiYWxlLCBVZ2FuZGEiLCJjb250YWN0IjoiKzk4ODc3NjU5ODciLCJlbWFpbCI6Im5hYnJyaWtrQHlhaG9vLmNvbSIsImV4cCI6MTYyMTEzMDI1NCwibmFtZSI6IlJpY2FyZG8gQmFua3MiLCJ1c2VyX2lkIjoxfQ.p3mrk5lGdfpECtSu749lDWLJI2wH1e8sxKq2SRjBnFA'
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
