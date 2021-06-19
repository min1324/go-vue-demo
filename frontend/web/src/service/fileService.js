import http from '@/plugins/axios';

const upload = (data) => http.post('/file/upload', data, { headers: { 'Content-Type': 'multipart/form-data' } });

export default upload;
