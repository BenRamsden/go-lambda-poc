import AuthProvider from "./components/auth/provider";
import { AuthenticatedGraphqlProvider } from "./components/gql/provider";
import { BrowserRouter, useRoutes } from "react-router-dom";

import Home from "./routes/home";
import Login from "./routes/login";
import ProtectedRoute from "./components/auth/protected-route";

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
    <AuthProvider>
      <AuthenticatedGraphqlProvider>
        <BrowserRouter>
          <Routes />
        </BrowserRouter>
      </AuthenticatedGraphqlProvider>
    </AuthProvider>
  );
}

export default App;

