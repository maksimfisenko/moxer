import App from "@/App";
import LoginPage from "@/features/auth/components/LoginPage";
import RegisterPage from "@/features/auth/components/RegisterPage";
import { createBrowserRouter } from "react-router";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
  },
  {
    path: "/login",
    element: <LoginPage />,
  },
]);

export default router;
