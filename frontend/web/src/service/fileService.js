import http from '@/plugins/http';

const upload = (data) => http.post('/file/upload', data, { headers: { 'Content-Type': 'multipart/form-data' } });

export default upload;
