import { useAuth0 } from "@auth0/auth0-react";
import { Button } from "@/components/ui/button";

const LoginButton = () => {
  const { loginWithRedirect } = useAuth0();

  return (
    <Button
      onClick={() => loginWithRedirect()}
      variant={"outline"}
      size={"default"}
    >
      Log In
    </Button>
  );
};

export default LoginButton;

