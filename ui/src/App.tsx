import AuthProvider from "./components/auth/provider";
import { AuthenticatedGraphqlProvider } from "./components/gql/provider";
import { BrowserRouter, useRoutes } from "react-router-dom";
import { ThemeProvider } from "@/components/theme-provider";

import Home from "./routes/home";
import Login from "./routes/login";
import ProtectedRoute from "./components/auth/protected-route";
import { Toaster } from "sonner";

const Routes = () => {
  const Router = useRoutes([
    {
      path: "/login",
      element: <Login />,
    },
    {
      path: "/",
      element: (
        <ProtectedRoute>
          <Home />
        </ProtectedRoute>
      ),
    },
  ]);

  return Router;
};

function App() {
  return (
    <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
      <AuthProvider>
        <AuthenticatedGraphqlProvider>
          <BrowserRouter>
            <Routes />
          </BrowserRouter>
          <Toaster />
        </AuthenticatedGraphqlProvider>
      </AuthProvider>
    </ThemeProvider>
  );
}

export default App;

