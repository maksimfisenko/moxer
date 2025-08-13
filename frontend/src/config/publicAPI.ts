import axios from "axios";

const publicAPI = axios.create({
  baseURL: "http://localhost:8080/api/v1/public",
});

export default publicAPI;
