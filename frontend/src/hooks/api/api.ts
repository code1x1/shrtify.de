import axios from "axios";

const api = axios.create({
    baseURL: import.meta.env.PUBLIC_API_BASE_URL,
    timeout: 2000,
});

api.interceptors.response.use(
    function (response) {
        return response;
    },
    function (error) {
        return Promise.resolve(error.response);
    }
);

export default api;
