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
  code: Scalars['String'];
  id: Scalars['ID'];
  image_logo: Scalars['String'];
  name: Scalars['String'];
};

export type Cart = {
  __typename?: 'Cart';
  id: Scalars['ID'];
  productItem: ProductItem;
  productItemId: Scalars['ID'];
  quantity: Scalars['Int'];
  user: User;
  userId: Scalars['ID'];
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
  addItemToCart?: Maybe<Array<Maybe<Cart>>>;
  createAdminUser: User;
  createUserBilling: UserBilling;
};


export type MutationAddItemToCartArgs = {
  itemId: Scalars['ID'];
  quantity?: InputMaybe<Scalars['Int']>;
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
  name?: Maybe<Scalars['String']>;
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

export type ProductItemPaginate = {
  __typename?: 'ProductItemPaginate';
  pagination?: Maybe<Pagination>;
  productItems?: Maybe<Array<Maybe<ProductItem>>>;
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

export type ProductSearchInput = {
  brand?: InputMaybe<Scalars['String']>;
  code?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  specifications?: InputMaybe<Array<InputMaybe<ProductSpecificationInput>>>;
};

export type ProductSpecification = {
  __typename?: 'ProductSpecification';
  code: Scalars['String'];
  id: Scalars['ID'];
  mandatory?: Maybe<Scalars['Boolean']>;
  name: Scalars['String'];
  productCategory?: Maybe<ProductCategory>;
  productCategoryID: Scalars['ID'];
  searchable?: Maybe<Scalars['Boolean']>;
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
  brands?: Maybe<Array<Maybe<Brand>>>;
  carts?: Maybe<Array<Maybe<Cart>>>;
  currencies?: Maybe<Array<Maybe<Currency>>>;
  currency?: Maybe<Currency>;
  legalEntityTypes?: Maybe<Array<Maybe<LegalEntityType>>>;
  productCategories?: Maybe<Array<Maybe<ProductCategory>>>;
  productItem?: Maybe<ProductItem>;
  productItems?: Maybe<ProductItemPaginate>;
  products?: Maybe<ProductPaginate>;
  productsItemsByCode?: Maybe<ProductItem>;
  searchBrands?: Maybe<Array<Maybe<Brand>>>;
  taxRates?: Maybe<Array<Maybe<Tax>>>;
  user?: Maybe<User>;
  users?: Maybe<Array<Maybe<User>>>;
};


export type QueryCurrencyArgs = {
  id: Scalars['ID'];
};


export type QueryProductItemArgs = {
  id: Scalars['ID'];
};


export type QueryProductItemsArgs = {
  pagination?: InputMaybe<PaginationInput>;
  productSearchInput?: InputMaybe<ProductSearchInput>;
};


export type QueryProductsArgs = {
  pagination?: InputMaybe<PaginationInput>;
  productSearchInput?: InputMaybe<ProductSearchInput>;
};


export type QueryProductsItemsByCodeArgs = {
  code: Scalars['String'];
};


export type QuerySearchBrandsArgs = {
  name: Scalars['String'];
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
  confirmed: Scalars['Boolean'];
  email: Scalars['String'];
  id: Scalars['ID'];
  name?: Maybe<Scalars['String']>;
  surname?: Maybe<Scalars['String']>;
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

export type AddCartMutationVariables = Exact<{
  itemId: Scalars['ID'];
  quantity?: InputMaybe<Scalars['Int']>;
}>;


export type AddCartMutation = { __typename?: 'Mutation', addItemToCart?: Array<{ __typename?: 'Cart', id: string, quantity: number, productItem: { __typename?: 'ProductItem', id: string, price: Array<{ __typename?: 'ProductPrice', value: number, currency: { __typename?: 'Currency', iso_code: string } } | null>, product: { __typename?: 'Product', name?: string | null, code: string } } } | null> | null };

export type CreateUserBillingMutationVariables = Exact<{
  input: CreateUserBilling;
}>;


export type CreateUserBillingMutation = { __typename?: 'Mutation', createUserBilling: { __typename?: 'UserBilling', id: string, name: string, surname: string } };

export type SearchBrandsQueryVariables = Exact<{
  name: Scalars['String'];
}>;


export type SearchBrandsQuery = { __typename?: 'Query', searchBrands?: Array<{ __typename?: 'Brand', id: string, name: string, code: string } | null> | null };

export type UserCartsQueryVariables = Exact<{ [key: string]: never; }>;


export type UserCartsQuery = { __typename?: 'Query', carts?: Array<{ __typename?: 'Cart', id: string, quantity: number, productItem: { __typename?: 'ProductItem', id: string, price: Array<{ __typename?: 'ProductPrice', value: number, currency: { __typename?: 'Currency', iso_code: string } } | null>, product: { __typename?: 'Product', name?: string | null, code: string } } } | null> | null };

export type GetLegalEntityTypesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetLegalEntityTypesQuery = { __typename?: 'Query', legalEntityTypes?: Array<{ __typename?: 'LegalEntityType', id: string, name: string, isPerson: boolean } | null> | null };

export type GetAllCategoriesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllCategoriesQuery = { __typename?: 'Query', productCategories?: Array<{ __typename?: 'ProductCategory', id: string, name: string, code: string, specifications: Array<{ __typename?: 'ProductSpecification', id: string, code: string, name: string, type: string, mandatory?: boolean | null, searchable?: boolean | null } | null> } | null> | null };

export type SearchQueryVariables = Exact<{
  limit: Scalars['Int'];
  offset: Scalars['Int'];
  searchInput?: InputMaybe<ProductSearchInput>;
}>;


export type SearchQuery = { __typename?: 'Query', productItems?: { __typename?: 'ProductItemPaginate', pagination?: { __typename?: 'Pagination', limit?: number | null, totals?: number | null, offset?: number | null } | null, productItems?: Array<{ __typename?: 'ProductItem', id: string, price: Array<{ __typename?: 'ProductPrice', value: number, currency: { __typename?: 'Currency', iso_code: string, symbol: string } } | null>, product: { __typename?: 'Product', id: string, name?: string | null, code: string, brand: { __typename?: 'Brand', name: string, code: string }, productSpecificationValues: Array<{ __typename?: 'ProductSpecificationValue', value: string, specification: { __typename?: 'ProductSpecification', code: string } } | null> } } | null> | null } | null };

export type ProductItemQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type ProductItemQuery = { __typename?: 'Query', productItem?: { __typename?: 'ProductItem', id: string, price: Array<{ __typename?: 'ProductPrice', value: number, currency: { __typename?: 'Currency', iso_code: string } } | null>, product: { __typename?: 'Product', code: string, name?: string | null, brand: { __typename?: 'Brand', name: string, code: string }, productSpecificationValues: Array<{ __typename?: 'ProductSpecificationValue', value: string, specification: { __typename?: 'ProductSpecification', code: string, name: string } } | null> } } | null };


export const AddCartDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addCart"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"itemId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"quantity"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addItemToCart"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"itemId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"itemId"}}},{"kind":"Argument","name":{"kind":"Name","value":"quantity"},"value":{"kind":"Variable","name":{"kind":"Name","value":"quantity"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"price"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}}]}}]}}]}}]}}]} as unknown as DocumentNode<AddCartMutation, AddCartMutationVariables>;
export const CreateUserBillingDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUserBilling"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateUserBilling"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUserBilling"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"billingInput"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}}]}}]}}]} as unknown as DocumentNode<CreateUserBillingMutation, CreateUserBillingMutationVariables>;
export const SearchBrandsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SearchBrands"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"searchBrands"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}}]}}]}}]} as unknown as DocumentNode<SearchBrandsQuery, SearchBrandsQueryVariables>;
export const UserCartsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"userCarts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"carts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"price"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}}]}}]}}]}}]}}]} as unknown as DocumentNode<UserCartsQuery, UserCartsQueryVariables>;
export const GetLegalEntityTypesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetLegalEntityTypes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"legalEntityTypes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"isPerson"}}]}}]}}]} as unknown as DocumentNode<GetLegalEntityTypesQuery, GetLegalEntityTypesQueryVariables>;
export const GetAllCategoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getAllCategories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productCategories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"specifications"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"mandatory"}},{"kind":"Field","name":{"kind":"Name","value":"searchable"}}]}}]}}]}}]} as unknown as DocumentNode<GetAllCategoriesQuery, GetAllCategoriesQueryVariables>;
export const SearchDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"search"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"limit"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"offset"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"searchInput"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"ProductSearchInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productItems"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"limit"},"value":{"kind":"Variable","name":{"kind":"Name","value":"limit"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"offset"},"value":{"kind":"Variable","name":{"kind":"Name","value":"offset"}}}]}},{"kind":"Argument","name":{"kind":"Name","value":"productSearchInput"},"value":{"kind":"Variable","name":{"kind":"Name","value":"searchInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pagination"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"totals"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItems"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"price"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productSpecificationValues"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"specification"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"value"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<SearchQuery, SearchQueryVariables>;
export const ProductItemDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"productItem"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productItem"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"price"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"value"}}]}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productSpecificationValues"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"specification"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ProductItemQuery, ProductItemQueryVariables>;