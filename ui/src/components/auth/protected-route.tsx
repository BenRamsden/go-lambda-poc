import { useAuth0 } from "@auth0/auth0-react";
import { PropsWithChildren, useEffect } from "react";
import { useNavigate } from "react-router-dom";

const ProtectedRoute = ({ children }: PropsWithChildren) => {
  const navigate = useNavigate();
  const { isAuthenticated, isLoading } = useAuth0();

  useEffect(() => {
    if (!isLoading && !isAuthenticated) {
      navigate("/login");
    }
  }, [isAuthenticated, isLoading]);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated) {
    return <div>Unauthorized</div>;
  }

  return <>{children}</>;
};

export default ProtectedRoute;

