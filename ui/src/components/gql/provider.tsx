import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
  ApolloLink,
  HttpLink,
} from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { useAuth0 } from "@auth0/auth0-react";
import { PropsWithChildren, useMemo } from "react";

const API_AUDIENCE: string = import.meta.env.VITE_API_AUDIENCE;

const CreateGraphqlClient = (link: ApolloLink) => {
  const httpLink = new HttpLink({
    uri: "http://localhost:4000/graphql",
  });

  return new ApolloClient({
    link: link.concat(httpLink),
    cache: new InMemoryCache(),
  });
};

export const GraphqlProvider = ({
  children,
  links,
}: PropsWithChildren<{
  links?: ApolloLink[];
}>) => {
  const client = CreateGraphqlClient(ApolloLink.from(links || []));

  return <ApolloProvider client={client}>{children}</ApolloProvider>;
};

export const AuthenticatedGraphqlProvider = ({
  children,
}: PropsWithChildren) => {
  const { getAccessTokenSilently } = useAuth0();

  const links = useMemo(() => {
    const authedLink = setContext(async (_, { headers }) => {
      const token = await getAccessTokenSilently({
        authorizationParams: {
          audience: API_AUDIENCE,
        },
      });

      return new ApolloLink((operation, forward) => {
        operation.setContext({
          headers: {
            ...headers,
            Authorization: `Bearer ${token}`,
          },
        });

        return forward(operation);
      });
    });

    return [authedLink];
  }, [getAccessTokenSilently]);

  return <GraphqlProvider links={links}>{children}</GraphqlProvider>;
};

