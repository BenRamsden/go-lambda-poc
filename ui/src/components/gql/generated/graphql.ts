/* eslint-disable */
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
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
  URI: Scalars['String']['output'];
  UpdatedAt: Scalars['Time']['output'];
};

export type Header = {
  __typename?: 'Header';
  Key: Scalars['String']['output'];
  Value: Scalars['String']['output'];
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
  deleteAsset: Scalars['Boolean']['output'];
  empty?: Maybe<Scalars['Boolean']['output']>;
  updateAsset: Asset;
  uploadAsset: PresignedUrl;
};


export type MutationCreateAssetArgs = {
  input: NewAsset;
};


export type MutationDeleteAssetArgs = {
  id: Scalars['ID']['input'];
};


export type MutationUpdateAssetArgs = {
  input: UpdateAsset;
};


export type MutationUploadAssetArgs = {
  id: Scalars['ID']['input'];
};

export type NewAsset = {
  Description: Scalars['String']['input'];
  Name: Scalars['String']['input'];
};

export type Pagination = {
  count: Scalars['Int']['input'];
  page: Scalars['Int']['input'];
};

export type PresignedUrl = {
  __typename?: 'PresignedURL';
  Fields: Array<Header>;
  URL: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  asset?: Maybe<Asset>;
  assets: Array<Asset>;
  me: Me;
};


export type QueryAssetArgs = {
  id: Scalars['ID']['input'];
};


export type QueryAssetsArgs = {
  pagination?: InputMaybe<Pagination>;
};

export type UpdateAsset = {
  Description?: InputMaybe<Scalars['String']['input']>;
  ID: Scalars['ID']['input'];
  Name?: InputMaybe<Scalars['String']['input']>;
};
