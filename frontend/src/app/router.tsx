import LoginPage from "@/features/auth/components/LoginPage";
import RegisterPage from "@/features/auth/components/RegisterPage";
import MainPage from "@/features/templates/components/MainPage";
import { createBrowserRouter } from "react-router";

const router = createBrowserRouter([
  {
    path: "/",
    element: <MainPage />,
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
