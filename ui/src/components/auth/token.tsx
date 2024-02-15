import { useAuth0 } from "@auth0/auth0-react";
import { useEffect } from "react";

const Token = () => {
  const { isAuthenticated } = useAuth0();

  useEffect(() => {
    // Force rerender to get the token
  }, [isAuthenticated]);

  return (
    <pre className="overflow-x-auto">
      {JSON.stringify(
        {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
        null,
        2
      )}
    </pre>
  );
};

export default Token;

