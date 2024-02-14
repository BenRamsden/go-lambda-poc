import { Auth0Provider } from "@auth0/auth0-react";
import { PropsWithChildren } from "react";

const AUTH0_DOMAIN = import.meta.env.VITE_AUTH0_DOMAIN;
const AUTH0_CLIENT_ID = import.meta.env.VITE_AUTH0_CLIENT_ID;

const AuthProvider = ({ children }: PropsWithChildren) => {
  return (
    <Auth0Provider
      domain={AUTH0_DOMAIN}
      clientId={AUTH0_CLIENT_ID}
      authorizationParams={{
        redirect_uri: "http://localhost:3000/",
      }}
    >
      {children}
    </Auth0Provider>
  );
};

export default AuthProvider;

