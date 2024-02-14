import { Auth0Provider, useAuth0 } from "@auth0/auth0-react";
import { useEffect, useState } from "react";

const AUTH0_DOMAIN = import.meta.env.VITE_AUTH0_DOMAIN;
const AUTH0_CLIENT_ID = import.meta.env.VITE_AUTH0_CLIENT_ID;
const API_AUDIENCE: string = import.meta.env.VITE_API_AUDIENCE;

const LoginButton = () => {
  const { loginWithRedirect } = useAuth0();

  return <button onClick={() => loginWithRedirect()}>Log In</button>;
};

const LogoutButton = () => {
  const { logout } = useAuth0();
  return (
    <button
      onClick={() =>
        logout({ logoutParams: { returnTo: window.location.origin } })
      }
    >
      Log Out
    </button>
  );
};

const Token = () => {
  const { getAccessTokenSilently } = useAuth0();
  const [token, setToken] = useState<{ Authorization: string } | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const token = await getAccessTokenSilently({
          authorizationParams: {
            audience: API_AUDIENCE,
          },
        });
        setToken({
          Authorization: `Bearer ${token}`,
        });
      } catch (e) {
        console.error(e);
      }
    })();
  }, []);

  return (
    <div>
      <h1>Token</h1>
      <pre>{JSON.stringify(token, undefined, 2)}</pre>
    </div>
  );
};

const LoggedIn = () => {
  const { user } = useAuth0();
  return (
    <div>
      <h1>Hello {user && user.name}</h1>
      <Token />
      <LogoutButton />
    </div>
  );
};

const LoggedOut = () => {
  return (
    <div>
      <h1>Hello! Please log in.</h1>
      <LoginButton />
    </div>
  );
};

const Main = () => {
  const { isAuthenticated, isLoading } = useAuth0();

  console.log(isLoading);

  if (isLoading) {
    return <div>Loading...</div>;
  }

  return <>{isAuthenticated ? <LoggedIn /> : <LoggedOut />}</>;
};

function App() {
  return (
    <>
      <Auth0Provider
        domain={AUTH0_DOMAIN}
        clientId={AUTH0_CLIENT_ID}
        authorizationParams={{
          redirect_uri: "http://localhost:3000/",
        }}
      >
        <Main />
      </Auth0Provider>
    </>
  );
}

export default App;

