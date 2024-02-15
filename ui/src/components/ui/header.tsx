import { useAuth0 } from "@auth0/auth0-react";
import { PropsWithChildren } from "react";
import LogoutButton from "../auth/logout-button";
import LoginButton from "../auth/login-button";

const LoginFlow = () => {
  const { isAuthenticated, isLoading } = useAuth0();

  if (isLoading) {
    return <></>;
  }

  if (isAuthenticated) {
    return <LogoutButton />;
  }

  return <LoginButton />;
};

const Header = ({}: PropsWithChildren) => {
  return (
    <header className="flex flex-row justify-between align-middle content-center p-4 bg-slate-200">
      <h1 className="self-center">Go POC</h1>
      <LoginFlow />
    </header>
  );
};

export default Header;

