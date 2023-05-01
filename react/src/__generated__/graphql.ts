/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
};

export type Brand = {
  __typename?: 'Brand';
  id: Scalars['ID'];
  image_logo: Scalars['String'];
  name: Scalars['String'];
};

export type CreateAdminUserInput = {
  email: Scalars['String'];
  name: Scalars['String'];
  password: Scalars['String'];
  surname: Scalars['String'];
};

export type CreateUserBilling = {
  addressLine1: Scalars['String'];
  addressLine2?: InputMaybe<Scalars['String']>;
  cap: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  fiscalCode?: InputMaybe<Scalars['String']>;
  iban: Scalars['String'];
  legalEntityTypeId: Scalars['ID'];
  name: Scalars['String'];
  province: Scalars['String'];
  surname?: InputMaybe<Scalars['String']>;
  vatNumber: Scalars['String'];
};

export type Currency = {
  __typename?: 'Currency';
  id: Scalars['ID'];
  iso_code: Scalars['String'];
  name: Scalars['String'];
  symbol: Scalars['String'];
  tag: Scalars['String'];
};

export type LegalEntityType = {
  __typename?: 'LegalEntityType';
  id: Scalars['ID'];
  isPerson: Scalars['Boolean'];
  name: Scalars['String'];
};

export type Mutation = {
  __typename?: 'Mutation';
  createAdminUser: User;
  createUserBilling: UserBilling;
};


export type MutationCreateAdminUserArgs = {
  userInput: CreateAdminUserInput;
};


export type MutationCreateUserBillingArgs = {
  billingInput: CreateUserBilling;
};

export type Pagination = {
  __typename?: 'Pagination';
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  totals?: Maybe<Scalars['Int']>;
};

export type PaginationInput = {
  limit: Scalars['Int'];
  offset: Scalars['Int'];
};

export type Product = {
  __typename?: 'Product';
  brand: Brand;
  brandID: Scalars['ID'];
  category: ProductCategory;
  code: Scalars['String'];
  id: Scalars['ID'];
  productCategoryID: Scalars['ID'];
  productSpecificationValues: Array<Maybe<ProductSpecificationValue>>;
  vehicleType: VehicleType;
  vehicleTypeID: Scalars['ID'];
};

export type ProductCategory = {
  __typename?: 'ProductCategory';
  code: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
  specifications: Array<Maybe<ProductSpecification>>;
};

export type ProductItem = {
  __typename?: 'ProductItem';
  id: Scalars['ID'];
  price: Array<Maybe<ProductPrice>>;
  product: Product;
  productID: Scalars['ID'];
  supplier: Supplier;
  supplierID: Scalars['ID'];
  supplier_price: Scalars['Float'];
  supplier_quantity: Scalars['Int'];
};

export type ProductPaginate = {
  __typename?: 'ProductPaginate';
  pagination?: Maybe<Pagination>;
  products?: Maybe<Array<Maybe<Product>>>;
};

export type ProductPrice = {
  __typename?: 'ProductPrice';
  currency: Currency;
  id: Scalars['ID'];
  value: Scalars['Float'];
};

export type ProductSpecification = {
  __typename?: 'ProductSpecification';
  code: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
  productCategory?: Maybe<ProductCategory>;
  productCategoryID: Scalars['ID'];
  type: Scalars['String'];
};

export type ProductSpecificationInput = {
  code: Scalars['String'];
  value: Scalars['String'];
};

export type ProductSpecificationValue = {
  __typename?: 'ProductSpecificationValue';
  id: Scalars['ID'];
  specification: ProductSpecification;
  specificationId: Scalars['ID'];
  value: Scalars['String'];
};

export type Query = {
  __typename?: 'Query';
  currencies?: Maybe<Array<Maybe<Currency>>>;
  currency?: Maybe<Currency>;
  legalEntityTypes?: Maybe<Array<Maybe<LegalEntityType>>>;
  productItems?: Maybe<Array<Maybe<ProductItem>>>;
  products?: Maybe<ProductPaginate>;
  productsItemsByCode?: Maybe<ProductItem>;
  taxRates?: Maybe<Array<Maybe<Tax>>>;
  user?: Maybe<User>;
  users?: Maybe<Array<Maybe<User>>>;
};


export type QueryCurrencyArgs = {
  id: Scalars['ID'];
};


export type QueryProductItemsArgs = {
  input: Array<InputMaybe<ProductSpecificationInput>>;
};


export type QueryProductsArgs = {
  pagination?: InputMaybe<PaginationInput>;
};


export type QueryProductsItemsByCodeArgs = {
  code: Scalars['String'];
};


export type QueryUserArgs = {
  ID: Scalars['ID'];
};

export type Supplier = {
  __typename?: 'Supplier';
  code: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
};

export type Tax = {
  __typename?: 'Tax';
  id: Scalars['ID'];
  markupPercentage: Scalars['Float'];
  name: Scalars['String'];
};

export type User = {
  __typename?: 'User';
  email: Scalars['String'];
  id: Scalars['ID'];
  userBilling: UserBilling;
  userRole: UserRole;
  username?: Maybe<Scalars['String']>;
};

export type UserBilling = {
  __typename?: 'UserBilling';
  addressLine1: Scalars['String'];
  addressLine2?: Maybe<Scalars['String']>;
  cap: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  fiscalCode: Scalars['String'];
  id: Scalars['ID'];
  legalEntityType: LegalEntityType;
  legalEntityTypeID?: Maybe<Scalars['ID']>;
  name: Scalars['String'];
  province?: Maybe<Scalars['String']>;
  surname: Scalars['String'];
  taxRate: Tax;
  user: User;
  userID?: Maybe<Scalars['ID']>;
  vatNumber: Scalars['String'];
};

export type UserRole = {
  __typename?: 'UserRole';
  id: Scalars['ID'];
  isAdmin?: Maybe<Scalars['Boolean']>;
  name: Scalars['String'];
  roleCode: Scalars['String'];
};

export type VehicleType = {
  __typename?: 'VehicleType';
  ID: Scalars['ID'];
  code: Scalars['String'];
  name: Scalars['String'];
};

export type CreateUserBillingMutationVariables = Exact<{
  input: CreateUserBilling;
}>;


export type CreateUserBillingMutation = { __typename?: 'Mutation', createUserBilling: { __typename?: 'UserBilling', id: string, name: string, surname: string } };

export type GetLegalEntityTypesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetLegalEntityTypesQuery = { __typename?: 'Query', legalEntityTypes?: Array<{ __typename?: 'LegalEntityType', id: string, name: string, isPerson: boolean } | null> | null };


export const CreateUserBillingDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUserBilling"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateUserBilling"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUserBilling"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"billingInput"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}}]}}]}}]} as unknown as DocumentNode<CreateUserBillingMutation, CreateUserBillingMutationVariables>;
export const GetLegalEntityTypesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetLegalEntityTypes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"legalEntityTypes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"isPerson"}}]}}]}}]} as unknown as DocumentNode<GetLegalEntityTypesQuery, GetLegalEntityTypesQueryVariables>;