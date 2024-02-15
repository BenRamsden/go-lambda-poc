import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  ApolloLink,
  HttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { useAuth0 } from "@auth0/auth0-react";
import { PropsWithChildren, useEffect } from "react";

const API_AUDIENCE: string = import.meta.env.VITE_API_AUDIENCE;

const CreateGraphqlClient = () => {
  const httpLink = new HttpLink({
    uri: "http://localhost:4000/graphql",
  });

  const authLink = setContext((_, { headers }) => {
    return {
      headers: {
        ...headers,
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    };
  });

  return new ApolloClient({
    link: ApolloLink.from([authLink, httpLink]),
    cache: new InMemoryCache(),
  });
};

export const GraphqlProvider = ({ children }: PropsWithChildren) => {
  const client = CreateGraphqlClient();

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};

export const AuthenticatedGraphqlProvider = ({
  children,
}: PropsWithChildren) => {
  const { getAccessTokenSilently } = useAuth0();

  useEffect(() => {
    const getToken = async () => {
      try {
        const token = await getAccessTokenSilently({
          authorizationParams: {
            audience: API_AUDIENCE,
          },
        });

        localStorage.setItem("token", token);
      } catch (e) {
        console.error(e);
      }
    };

    getToken();

    return () => {
      localStorage.removeItem("token");
    };
  }, [getAccessTokenSilently]);

  return <GraphqlProvider>{children}</GraphqlProvider>;
};

