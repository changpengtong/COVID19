import axios from "axios"

let http = axios.create({
    timeout: 100000
})

if (process.env.NODE_ENV == 'development') {
    http.defaults.baseURL = 'localhost:9999/COVID19/';}
else if (process.env.NODE_ENV == 'production') {
    http.defaults.baseURL = '/COVID19/';
}

export default http;
