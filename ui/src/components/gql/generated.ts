import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  Time: { input: any; output: any; }
};

export type Asset = {
  __typename?: 'Asset';
  CreatedAt: Scalars['Time']['output'];
  Description: Scalars['String']['output'];
  ID: Scalars['ID']['output'];
  Name: Scalars['String']['output'];
  Owner: Scalars['ID']['output'];
  URI: Scalars['String']['output'];
  UpdatedAt: Scalars['Time']['output'];
};

export type Me = {
  __typename?: 'Me';
  email: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createAsset: Asset;
  createPanic: Scalars['Boolean']['output'];
  empty?: Maybe<Scalars['Boolean']['output']>;
};


export type MutationCreateAssetArgs = {
  input: NewAsset;
};


export type MutationCreatePanicArgs = {
  message: Scalars['String']['input'];
};

export type NewAsset = {
  Description: Scalars['String']['input'];
  Name: Scalars['String']['input'];
  URI: Scalars['String']['input'];
};

export type Query = {
  __typename?: 'Query';
  assets: Array<Asset>;
  getPanic: Scalars['Boolean']['output'];
  me: Me;
};


export type QueryGetPanicArgs = {
  message: Scalars['String']['input'];
};

export type GetAssetsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAssetsQuery = { __typename?: 'Query', assets: Array<{ __typename?: 'Asset', ID: string, Owner: string, Name: string, Description: string, URI: string, CreatedAt: any, UpdatedAt: any }> };

export type CreateAssetMutationVariables = Exact<{
  input: NewAsset;
}>;


export type CreateAssetMutation = { __typename?: 'Mutation', createAsset: { __typename?: 'Asset', ID: string, Owner: string, Name: string, Description: string, URI: string, CreatedAt: any, UpdatedAt: any } };

export type GetPanicQueryVariables = Exact<{
  message: Scalars['String']['input'];
}>;


export type GetPanicQuery = { __typename?: 'Query', getPanic: boolean };

export type CreatePanicMutationVariables = Exact<{
  message: Scalars['String']['input'];
}>;


export type CreatePanicMutation = { __typename?: 'Mutation', createPanic: boolean };


export const GetAssetsDocument = gql`
    query getAssets {
  assets {
    ID
    Owner
    Name
    Description
    URI
    CreatedAt
    UpdatedAt
  }
}
    `;

/**
 * __useGetAssetsQuery__
 *
 * To run a query within a React component, call `useGetAssetsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAssetsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAssetsQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetAssetsQuery(baseOptions?: Apollo.QueryHookOptions<GetAssetsQuery, GetAssetsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAssetsQuery, GetAssetsQueryVariables>(GetAssetsDocument, options);
      }
export function useGetAssetsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAssetsQuery, GetAssetsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAssetsQuery, GetAssetsQueryVariables>(GetAssetsDocument, options);
        }
export function useGetAssetsSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetAssetsQuery, GetAssetsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetAssetsQuery, GetAssetsQueryVariables>(GetAssetsDocument, options);
        }
export type GetAssetsQueryHookResult = ReturnType<typeof useGetAssetsQuery>;
export type GetAssetsLazyQueryHookResult = ReturnType<typeof useGetAssetsLazyQuery>;
export type GetAssetsSuspenseQueryHookResult = ReturnType<typeof useGetAssetsSuspenseQuery>;
export type GetAssetsQueryResult = Apollo.QueryResult<GetAssetsQuery, GetAssetsQueryVariables>;
export const CreateAssetDocument = gql`
    mutation createAsset($input: NewAsset!) {
  createAsset(input: $input) {
    ID
    Owner
    Name
    Description
    URI
    CreatedAt
    UpdatedAt
  }
}
    `;
export type CreateAssetMutationFn = Apollo.MutationFunction<CreateAssetMutation, CreateAssetMutationVariables>;

/**
 * __useCreateAssetMutation__
 *
 * To run a mutation, you first call `useCreateAssetMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateAssetMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createAssetMutation, { data, loading, error }] = useCreateAssetMutation({
 *   variables: {
 *      input: // value for 'input'
 *   },
 * });
 */
export function useCreateAssetMutation(baseOptions?: Apollo.MutationHookOptions<CreateAssetMutation, CreateAssetMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateAssetMutation, CreateAssetMutationVariables>(CreateAssetDocument, options);
      }
export type CreateAssetMutationHookResult = ReturnType<typeof useCreateAssetMutation>;
export type CreateAssetMutationResult = Apollo.MutationResult<CreateAssetMutation>;
export type CreateAssetMutationOptions = Apollo.BaseMutationOptions<CreateAssetMutation, CreateAssetMutationVariables>;
export const GetPanicDocument = gql`
    query getPanic($message: String!) {
  getPanic(message: $message)
}
    `;

/**
 * __useGetPanicQuery__
 *
 * To run a query within a React component, call `useGetPanicQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetPanicQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetPanicQuery({
 *   variables: {
 *      message: // value for 'message'
 *   },
 * });
 */
export function useGetPanicQuery(baseOptions: Apollo.QueryHookOptions<GetPanicQuery, GetPanicQueryVariables> & ({ variables: GetPanicQueryVariables; skip?: boolean; } | { skip: boolean; }) ) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetPanicQuery, GetPanicQueryVariables>(GetPanicDocument, options);
      }
export function useGetPanicLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetPanicQuery, GetPanicQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetPanicQuery, GetPanicQueryVariables>(GetPanicDocument, options);
        }
export function useGetPanicSuspenseQuery(baseOptions?: Apollo.SuspenseQueryHookOptions<GetPanicQuery, GetPanicQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useSuspenseQuery<GetPanicQuery, GetPanicQueryVariables>(GetPanicDocument, options);
        }
export type GetPanicQueryHookResult = ReturnType<typeof useGetPanicQuery>;
export type GetPanicLazyQueryHookResult = ReturnType<typeof useGetPanicLazyQuery>;
export type GetPanicSuspenseQueryHookResult = ReturnType<typeof useGetPanicSuspenseQuery>;
export type GetPanicQueryResult = Apollo.QueryResult<GetPanicQuery, GetPanicQueryVariables>;
export const CreatePanicDocument = gql`
    mutation createPanic($message: String!) {
  createPanic(message: $message)
}
    `;
export type CreatePanicMutationFn = Apollo.MutationFunction<CreatePanicMutation, CreatePanicMutationVariables>;

/**
 * __useCreatePanicMutation__
 *
 * To run a mutation, you first call `useCreatePanicMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreatePanicMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createPanicMutation, { data, loading, error }] = useCreatePanicMutation({
 *   variables: {
 *      message: // value for 'message'
 *   },
 * });
 */
export function useCreatePanicMutation(baseOptions?: Apollo.MutationHookOptions<CreatePanicMutation, CreatePanicMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreatePanicMutation, CreatePanicMutationVariables>(CreatePanicDocument, options);
      }
export type CreatePanicMutationHookResult = ReturnType<typeof useCreatePanicMutation>;
export type CreatePanicMutationResult = Apollo.MutationResult<CreatePanicMutation>;
export type CreatePanicMutationOptions = Apollo.BaseMutationOptions<CreatePanicMutation, CreatePanicMutationVariables>;