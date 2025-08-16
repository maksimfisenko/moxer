import Content from "../components/Content";
import { useNavigate } from "react-router";
import { useQueryClient } from "@tanstack/react-query";
import LoadingSpinner from "@/shared/components/LoadingSpinner";
import { useGetCurrentUser } from "@/features/header/hooks/useGetCurrentUser";
import Header from "@/features/header/components/Header";
import MainLayout from "../layouts/MainLayout";

const MainPage = () => {
  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { data, isLoading, isError } = useGetCurrentUser();

  const handleLogout = () => {
    localStorage.removeItem("token");
    queryClient.removeQueries({ queryKey: ["me"] });
    navigate("/login");
  };

  if (isLoading) return <LoadingSpinner />;

  if (isError || !data) {
    handleLogout();
    return null;
  }

  return (
    <MainLayout>
      <Header user={data} onButtonClick={handleLogout} />
      <Content />
    </MainLayout>
  );
};

export default MainPage;
