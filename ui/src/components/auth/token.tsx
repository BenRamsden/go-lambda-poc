import { useAuth0 } from "@auth0/auth0-react";
import { useEffect, useState } from "react";

const Token = () => {
  const [auth0Token, setAuth0Token] = useState<{ Authorization: string }>({
    Authorization: "",
  });
  const { getAccessTokenSilently } = useAuth0();

  useEffect(() => {
    const getToken = async () => {
      try {
        const token = await getAccessTokenSilently();
        setAuth0Token({ Authorization: `Bearer ${token}` });
      } catch (e) {
        console.error(e);
      }
    };

    getToken();
  }, [getAccessTokenSilently]);

  return (
    <pre className="overflow-x-auto">{JSON.stringify(auth0Token, null, 2)}</pre>
  );
};

export default Token;

