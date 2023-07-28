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
  Timestamp: any;
};

export type Brand = {
  __typename?: 'Brand';
  code: Scalars['String'];
  id: Scalars['ID'];
  image_logo: Scalars['String'];
  name: Scalars['String'];
  quality?: Maybe<Scalars['Int']>;
};

export type Cart = {
  __typename?: 'Cart';
  id: Scalars['ID'];
  productItemPrice: ProductItemPrice;
  productItemPriceId: Scalars['ID'];
  quantity: Scalars['Int'];
  user: User;
  userId: Scalars['ID'];
};

export type CartResponse = {
  __typename?: 'CartResponse';
  items: Array<Maybe<Cart>>;
  totalPrice?: Maybe<TotalPrice>;
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
  sdiCode?: InputMaybe<Scalars['String']>;
  sdiPec?: InputMaybe<Scalars['String']>;
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
  addItemToCart?: Maybe<CartResponse>;
  createAdminUser: User;
  createUserAddress: Array<Maybe<UserAddress>>;
  createUserBilling: UserBilling;
  deleteUserAddress: Array<Maybe<UserAddress>>;
  editCart?: Maybe<CartResponse>;
  editUserAddress: Array<Maybe<UserAddress>>;
  emptyCart?: Maybe<CartResponse>;
  newOrder: Order;
  updateUserStatus: User;
};


export type MutationAddItemToCartArgs = {
  itemId: Scalars['ID'];
  quantity?: InputMaybe<Scalars['Int']>;
};


export type MutationCreateAdminUserArgs = {
  userInput: CreateAdminUserInput;
};


export type MutationCreateUserAddressArgs = {
  userAddress: UserAddressInput;
};


export type MutationCreateUserBillingArgs = {
  billingInput: CreateUserBilling;
};


export type MutationDeleteUserAddressArgs = {
  id: Scalars['ID'];
};


export type MutationEditCartArgs = {
  cartId: Scalars['ID'];
  quantity: Scalars['Int'];
};


export type MutationEditUserAddressArgs = {
  id: Scalars['ID'];
  userAddress: UserAddressInput;
};


export type MutationNewOrderArgs = {
  userAddressId: Scalars['ID'];
  userId: Scalars['ID'];
};


export type MutationUpdateUserStatusArgs = {
  confirmed?: InputMaybe<Scalars['Boolean']>;
  rejected?: InputMaybe<Scalars['Boolean']>;
  userID: Scalars['ID'];
};

export type Order = {
  __typename?: 'Order';
  addressLine1: Scalars['String'];
  addressLine2?: Maybe<Scalars['String']>;
  addressName: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  createdAt: Scalars['Timestamp'];
  currency: Currency;
  currencyID: Scalars['ID'];
  id: Scalars['ID'];
  orderRows: Array<Maybe<OrderRow>>;
  postalCode: Scalars['String'];
  priceAmount: Scalars['Float'];
  province: Scalars['String'];
  status: OrderStatus;
  tax: Tax;
  taxID: Scalars['ID'];
  userBilling: UserBilling;
  userBillingID: Scalars['ID'];
};

export type OrderFilterInput = {
  dateFrom?: InputMaybe<Scalars['String']>;
  dateTo?: InputMaybe<Scalars['String']>;
  number?: InputMaybe<Scalars['String']>;
};

export type OrderRow = {
  __typename?: 'OrderRow';
  amount: Scalars['Float'];
  deliveredAt?: Maybe<Scalars['Timestamp']>;
  id: Scalars['ID'];
  order: Order;
  orderID: Scalars['ID'];
  productItemPrice: ProductItemPrice;
  productItemPriceID: Scalars['ID'];
  quantity: Scalars['Int'];
  trackingNumber?: Maybe<Scalars['String']>;
};

export enum OrderStatus {
  Canceled = 'CANCELED',
  Confirmed = 'CONFIRMED',
  Delivered = 'DELIVERED',
  New = 'NEW',
  Rejected = 'REJECTED',
  Returned = 'RETURNED'
}

export type OrderingInput = {
  column: Scalars['String'];
  desc?: InputMaybe<Scalars['Boolean']>;
};

export type OrdersFilterInput = {
  dateFrom?: InputMaybe<Scalars['String']>;
  dateTo?: InputMaybe<Scalars['String']>;
  number?: InputMaybe<Scalars['String']>;
  status?: InputMaybe<OrderStatus>;
};

export type OrdersPaginator = {
  __typename?: 'OrdersPaginator';
  data: Array<Maybe<Order>>;
  pagination: Pagination;
};

export type Pagination = {
  __typename?: 'Pagination';
  currentPage?: Maybe<Scalars['Int']>;
  limit?: Maybe<Scalars['Int']>;
  offset?: Maybe<Scalars['Int']>;
  pageCount?: Maybe<Scalars['Int']>;
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
  eprelProductCode?: Maybe<Scalars['String']>;
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
  price: Array<Maybe<ProductItemPrice>>;
  product: Product;
  productID: Scalars['ID'];
  supplier: Supplier;
  supplierID: Scalars['ID'];
  supplierPrice: Scalars['Float'];
  supplierQuantity: Scalars['Int'];
};

export type ProductItemPaginate = {
  __typename?: 'ProductItemPaginate';
  pagination?: Maybe<Pagination>;
  productItems?: Maybe<Array<Maybe<ProductItem>>>;
};

export type ProductItemPrice = {
  __typename?: 'ProductItemPrice';
  currency: Currency;
  id: Scalars['ID'];
  productItem: ProductItem;
  productItemId: Scalars['ID'];
  value: Scalars['Float'];
};

export type ProductPaginate = {
  __typename?: 'ProductPaginate';
  pagination?: Maybe<Pagination>;
  products?: Maybe<Array<Maybe<Product>>>;
};

export type ProductSearchInput = {
  brand?: InputMaybe<Scalars['String']>;
  code?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  specifications?: InputMaybe<Array<InputMaybe<ProductSpecificationInput>>>;
  vehicleCode?: InputMaybe<Scalars['String']>;
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
  values: Array<Maybe<ProductSpecificationValue>>;
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
  allOrders: OrdersPaginator;
  brands?: Maybe<Array<Maybe<Brand>>>;
  carts?: Maybe<CartResponse>;
  currencies?: Maybe<Array<Maybe<Currency>>>;
  currency?: Maybe<Currency>;
  legalEntityTypes?: Maybe<Array<Maybe<LegalEntityType>>>;
  order: Order;
  productCategories?: Maybe<Array<Maybe<ProductCategory>>>;
  productItem?: Maybe<ProductItem>;
  productItems?: Maybe<ProductItemPaginate>;
  products?: Maybe<ProductPaginate>;
  productsItemsByCode?: Maybe<ProductItem>;
  searchBrands?: Maybe<Array<Maybe<Brand>>>;
  searchSpecificationValue?: Maybe<Array<Maybe<ProductSpecificationValue>>>;
  specifications?: Maybe<Array<Maybe<ProductSpecification>>>;
  taxRates?: Maybe<Array<Maybe<Tax>>>;
  user?: Maybe<User>;
  userAddress?: Maybe<Array<Maybe<UserAddress>>>;
  userBilling?: Maybe<UserBilling>;
  userOrders?: Maybe<OrdersPaginator>;
  users?: Maybe<UserPaginator>;
};


export type QueryAllOrdersArgs = {
  filter?: InputMaybe<OrdersFilterInput>;
  ordering?: InputMaybe<Array<InputMaybe<OrderingInput>>>;
  pagination?: InputMaybe<PaginationInput>;
};


export type QueryCurrencyArgs = {
  id: Scalars['ID'];
};


export type QueryOrderArgs = {
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


export type QuerySearchSpecificationValueArgs = {
  code: Scalars['String'];
  value?: InputMaybe<Scalars['String']>;
  vehicleCode?: InputMaybe<Scalars['String']>;
};


export type QueryUserArgs = {
  ID: Scalars['ID'];
};


export type QueryUserAddressArgs = {
  userId: Scalars['ID'];
};


export type QueryUserBillingArgs = {
  userId: Scalars['ID'];
};


export type QueryUserOrdersArgs = {
  filter?: InputMaybe<OrderFilterInput>;
  ordering?: InputMaybe<Array<InputMaybe<OrderingInput>>>;
  pagination?: InputMaybe<PaginationInput>;
  userId: Scalars['ID'];
};


export type QueryUsersArgs = {
  filter?: InputMaybe<UserFilterInput>;
  pagination?: InputMaybe<PaginationInput>;
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

export type TotalPrice = {
  __typename?: 'TotalPrice';
  currency?: Maybe<Currency>;
  value: Scalars['Float'];
};

export type User = {
  __typename?: 'User';
  confirmed: Scalars['Boolean'];
  email: Scalars['String'];
  id: Scalars['ID'];
  name?: Maybe<Scalars['String']>;
  rejected: Scalars['Boolean'];
  surname?: Maybe<Scalars['String']>;
  userBilling?: Maybe<UserBilling>;
  userRole: UserRole;
  userRoleId: Scalars['ID'];
  username?: Maybe<Scalars['String']>;
};

export type UserAddress = {
  __typename?: 'UserAddress';
  ID: Scalars['ID'];
  addressLine1: Scalars['String'];
  addressLine2?: Maybe<Scalars['String']>;
  addressName: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  isDefault: Scalars['Boolean'];
  postalCode: Scalars['String'];
  province: Scalars['String'];
  user: User;
  userID: Scalars['ID'];
};

export type UserAddressInput = {
  IsDefault: Scalars['Boolean'];
  addressLine1: Scalars['String'];
  addressLine2?: InputMaybe<Scalars['String']>;
  addressName: Scalars['String'];
  city: Scalars['String'];
  country: Scalars['String'];
  postalCode: Scalars['String'];
  province: Scalars['String'];
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
  sdiCode?: Maybe<Scalars['String']>;
  sdiPec?: Maybe<Scalars['String']>;
  surname: Scalars['String'];
  taxRate: Tax;
  user: User;
  userID?: Maybe<Scalars['ID']>;
  vatNumber: Scalars['String'];
};

export type UserFilterInput = {
  confirmed?: InputMaybe<Scalars['Boolean']>;
  email?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
};

export type UserPaginator = {
  __typename?: 'UserPaginator';
  data: Array<Maybe<User>>;
  pagination: Pagination;
};

export type UserRole = {
  __typename?: 'UserRole';
  id: Scalars['ID'];
  isAdmin: Scalars['Boolean'];
  name: Scalars['String'];
  roleCode: Scalars['String'];
};

export type VehicleType = {
  __typename?: 'VehicleType';
  ID: Scalars['ID'];
  code: Scalars['String'];
  name: Scalars['String'];
};

export type CartItemsFragment = { __typename?: 'CartResponse', totalPrice?: { __typename?: 'TotalPrice', value: number, currency?: { __typename?: 'Currency', iso_code: string, symbol: string, name: string } | null } | null, items: Array<{ __typename?: 'Cart', id: string, quantity: number, productItemPrice: { __typename?: 'ProductItemPrice', value: number, currency: { __typename?: 'Currency', iso_code: string }, productItem: { __typename?: 'ProductItem', id: string, product: { __typename?: 'Product', name?: string | null, code: string, eprelProductCode?: string | null, brand: { __typename?: 'Brand', name: string } } } } } | null> } & { ' $fragmentName'?: 'CartItemsFragment' };

export type PaginationInfoFragment = { __typename?: 'Pagination', limit?: number | null, totals?: number | null, offset?: number | null, currentPage?: number | null, pageCount?: number | null } & { ' $fragmentName'?: 'PaginationInfoFragment' };

export type UserAddressCollectionFragment = { __typename?: 'UserAddress', ID: string, addressName: string, isDefault: boolean, addressLine1: string, addressLine2?: string | null, city: string, country: string, postalCode: string, province: string } & { ' $fragmentName'?: 'UserAddressCollectionFragment' };

export type AddCartMutationVariables = Exact<{
  itemId: Scalars['ID'];
  quantity?: InputMaybe<Scalars['Int']>;
}>;


export type AddCartMutation = { __typename?: 'Mutation', addItemToCart?: (
    { __typename?: 'CartResponse' }
    & { ' $fragmentRefs'?: { 'CartItemsFragment': CartItemsFragment } }
  ) | null };

export type EditCartMutationVariables = Exact<{
  cartId: Scalars['ID'];
  quantity: Scalars['Int'];
}>;


export type EditCartMutation = { __typename?: 'Mutation', editCart?: (
    { __typename?: 'CartResponse' }
    & { ' $fragmentRefs'?: { 'CartItemsFragment': CartItemsFragment } }
  ) | null };

export type EmptyMutationVariables = Exact<{ [key: string]: never; }>;


export type EmptyMutation = { __typename?: 'Mutation', emptyCart?: (
    { __typename?: 'CartResponse' }
    & { ' $fragmentRefs'?: { 'CartItemsFragment': CartItemsFragment } }
  ) | null };

export type CreateUserBillingMutationVariables = Exact<{
  input: CreateUserBilling;
}>;


export type CreateUserBillingMutation = { __typename?: 'Mutation', createUserBilling: { __typename?: 'UserBilling', id: string, name: string, surname: string } };

export type CreateNewOrderMutationVariables = Exact<{
  userId: Scalars['ID'];
  userAddressId: Scalars['ID'];
}>;


export type CreateNewOrderMutation = { __typename?: 'Mutation', newOrder: { __typename?: 'Order', id: string } };

export type AddAddressMutationVariables = Exact<{
  input: UserAddressInput;
}>;


export type AddAddressMutation = { __typename?: 'Mutation', createUserAddress: Array<(
    { __typename?: 'UserAddress' }
    & { ' $fragmentRefs'?: { 'UserAddressCollectionFragment': UserAddressCollectionFragment } }
  ) | null> };

export type EditAddressMutationVariables = Exact<{
  id: Scalars['ID'];
  input: UserAddressInput;
}>;


export type EditAddressMutation = { __typename?: 'Mutation', editUserAddress: Array<(
    { __typename?: 'UserAddress' }
    & { ' $fragmentRefs'?: { 'UserAddressCollectionFragment': UserAddressCollectionFragment } }
  ) | null> };

export type DeleteAddressMutationVariables = Exact<{
  id: Scalars['ID'];
}>;


export type DeleteAddressMutation = { __typename?: 'Mutation', deleteUserAddress: Array<(
    { __typename?: 'UserAddress' }
    & { ' $fragmentRefs'?: { 'UserAddressCollectionFragment': UserAddressCollectionFragment } }
  ) | null> };

export type ChangeUserStatusMutationVariables = Exact<{
  userID: Scalars['ID'];
  confirmed?: InputMaybe<Scalars['Boolean']>;
  rejected?: InputMaybe<Scalars['Boolean']>;
}>;


export type ChangeUserStatusMutation = { __typename?: 'Mutation', updateUserStatus: { __typename?: 'User', id: string } };

export type SearchBrandsQueryVariables = Exact<{
  name: Scalars['String'];
}>;


export type SearchBrandsQuery = { __typename?: 'Query', searchBrands?: Array<{ __typename?: 'Brand', id: string, name: string, code: string, quality?: number | null } | null> | null };

export type UserCartsQueryVariables = Exact<{ [key: string]: never; }>;


export type UserCartsQuery = { __typename?: 'Query', carts?: { __typename?: 'CartResponse', totalPrice?: { __typename?: 'TotalPrice', value: number, currency?: { __typename?: 'Currency', iso_code: string, symbol: string, name: string } | null } | null, items: Array<{ __typename?: 'Cart', id: string, quantity: number, productItemPrice: { __typename?: 'ProductItemPrice', value: number, currency: { __typename?: 'Currency', iso_code: string }, productItem: { __typename?: 'ProductItem', id: string, product: { __typename?: 'Product', name?: string | null, code: string, brand: { __typename?: 'Brand', name: string } } } } } | null> } | null };

export type GetLegalEntityTypesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetLegalEntityTypesQuery = { __typename?: 'Query', legalEntityTypes?: Array<{ __typename?: 'LegalEntityType', id: string, name: string, isPerson: boolean } | null> | null };

export type FetchOrderQueryVariables = Exact<{
  orderId: Scalars['ID'];
}>;


export type FetchOrderQuery = { __typename?: 'Query', order: { __typename?: 'Order', id: string, priceAmount: number, addressName: string, addressLine1: string, addressLine2?: string | null, city: string, country: string, province: string, postalCode: string, status: OrderStatus, createdAt: any, currency: { __typename?: 'Currency', iso_code: string, name: string }, userBilling: { __typename?: 'UserBilling', id: string, name: string, surname: string, vatNumber: string, fiscalCode: string, legalEntityType: { __typename?: 'LegalEntityType', name: string } }, orderRows: Array<{ __typename?: 'OrderRow', id: string, amount: number, quantity: number, productItemPrice: { __typename?: 'ProductItemPrice', value: number, currency: { __typename?: 'Currency', iso_code: string, symbol: string }, productItem: { __typename?: 'ProductItem', id: string, product: { __typename?: 'Product', id: string, name?: string | null } } } } | null> } };

export type FetchOrdersQueryVariables = Exact<{
  userId: Scalars['ID'];
  pagination?: InputMaybe<PaginationInput>;
  filter?: InputMaybe<OrderFilterInput>;
  ordering?: InputMaybe<Array<InputMaybe<OrderingInput>> | InputMaybe<OrderingInput>>;
}>;


export type FetchOrdersQuery = { __typename?: 'Query', userOrders?: { __typename?: 'OrdersPaginator', data: Array<{ __typename?: 'Order', id: string, status: OrderStatus, addressName: string, addressLine1: string, addressLine2?: string | null, city: string, province: string, postalCode: string, country: string, createdAt: any, priceAmount: number, currency: { __typename?: 'Currency', iso_code: string }, orderRows: Array<{ __typename?: 'OrderRow', amount: number, quantity: number } | null> } | null>, pagination: (
      { __typename?: 'Pagination' }
      & { ' $fragmentRefs'?: { 'PaginationInfoFragment': PaginationInfoFragment } }
    ) } | null };

export type AllOrdersQueryVariables = Exact<{
  filter?: InputMaybe<OrdersFilterInput>;
  pagination?: InputMaybe<PaginationInput>;
  ordering?: InputMaybe<Array<InputMaybe<OrderingInput>> | InputMaybe<OrderingInput>>;
}>;


export type AllOrdersQuery = { __typename?: 'Query', allOrders: { __typename?: 'OrdersPaginator', data: Array<{ __typename?: 'Order', id: string, priceAmount: number, currency: { __typename?: 'Currency', iso_code: string }, userBilling: { __typename?: 'UserBilling', name: string, surname: string, user: { __typename?: 'User', email: string } }, orderRows: Array<{ __typename?: 'OrderRow', id: string, amount: number } | null> } | null>, pagination: (
      { __typename?: 'Pagination' }
      & { ' $fragmentRefs'?: { 'PaginationInfoFragment': PaginationInfoFragment } }
    ) } };

export type GetAllCategoriesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAllCategoriesQuery = { __typename?: 'Query', productCategories?: Array<{ __typename?: 'ProductCategory', id: string, name: string, code: string, specifications: Array<{ __typename?: 'ProductSpecification', id: string, code: string, name: string, type: string, mandatory?: boolean | null, searchable?: boolean | null } | null> } | null> | null };

export type SearchQueryVariables = Exact<{
  limit: Scalars['Int'];
  offset: Scalars['Int'];
  searchInput?: InputMaybe<ProductSearchInput>;
}>;


export type SearchQuery = { __typename?: 'Query', productItems?: { __typename?: 'ProductItemPaginate', pagination?: { __typename?: 'Pagination', limit?: number | null, totals?: number | null, offset?: number | null, currentPage?: number | null, pageCount?: number | null } | null, productItems?: Array<{ __typename?: 'ProductItem', id: string, supplierQuantity: number, price: Array<{ __typename?: 'ProductItemPrice', value: number, currency: { __typename?: 'Currency', iso_code: string, symbol: string } } | null>, product: { __typename?: 'Product', id: string, name?: string | null, code: string, eprelProductCode?: string | null, brand: { __typename?: 'Brand', name: string, code: string, quality?: number | null }, vehicleType: { __typename?: 'VehicleType', name: string }, productSpecificationValues: Array<{ __typename?: 'ProductSpecificationValue', value: string, specification: { __typename?: 'ProductSpecification', code: string } } | null> } } | null> | null } | null };

export type ProductItemQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type ProductItemQuery = { __typename?: 'Query', productItem?: { __typename?: 'ProductItem', id: string, supplierQuantity: number, price: Array<{ __typename?: 'ProductItemPrice', value: number, currency: { __typename?: 'Currency', iso_code: string } } | null>, product: { __typename?: 'Product', code: string, name?: string | null, eprelProductCode?: string | null, brand: { __typename?: 'Brand', name: string, code: string, quality?: number | null }, vehicleType: { __typename?: 'VehicleType', name: string }, productSpecificationValues: Array<{ __typename?: 'ProductSpecificationValue', value: string, specification: { __typename?: 'ProductSpecification', code: string, name: string } } | null> } } | null };

export type SearchValuesQueryVariables = Exact<{
  code: Scalars['String'];
  value?: InputMaybe<Scalars['String']>;
  vehicleCode?: InputMaybe<Scalars['String']>;
}>;


export type SearchValuesQuery = { __typename?: 'Query', searchSpecificationValue?: Array<{ __typename?: 'ProductSpecificationValue', value: string } | null> | null };

export type FetchUserAddressesQueryVariables = Exact<{
  userId: Scalars['ID'];
}>;


export type FetchUserAddressesQuery = { __typename?: 'Query', userAddress?: Array<{ __typename?: 'UserAddress', ID: string, addressName: string, isDefault: boolean, addressLine1: string, addressLine2?: string | null, city: string, country: string, postalCode: string, province: string } | null> | null };

export type FetchUserQueryQueryVariables = Exact<{
  userId: Scalars['ID'];
}>;


export type FetchUserQueryQuery = { __typename?: 'Query', userBilling?: { __typename?: 'UserBilling', id: string, name: string, surname: string, addressLine1: string, addressLine2?: string | null, city: string, country: string, province?: string | null, cap: string, fiscalCode: string, vatNumber: string, sdiCode?: string | null, sdiPec?: string | null, legalEntityType: { __typename?: 'LegalEntityType', name: string }, user: { __typename?: 'User', id: string, name?: string | null, surname?: string | null, email: string } } | null };

export type FetchUserQueryVariables = Exact<{
  userId: Scalars['ID'];
}>;


export type FetchUserQuery = { __typename?: 'Query', user?: { __typename?: 'User', id: string, name?: string | null, surname?: string | null, email: string, confirmed: boolean, rejected: boolean, userRole: { __typename?: 'UserRole', isAdmin: boolean, name: string }, userBilling?: { __typename?: 'UserBilling', id: string, name: string, surname: string, addressLine1: string, addressLine2?: string | null, city: string, country: string, province?: string | null, cap: string, fiscalCode: string, vatNumber: string, sdiCode?: string | null, sdiPec?: string | null, legalEntityType: { __typename?: 'LegalEntityType', name: string } } | null } | null };

export type FetchAllUsersQueryVariables = Exact<{
  pagination?: InputMaybe<PaginationInput>;
  filter?: InputMaybe<UserFilterInput>;
}>;


export type FetchAllUsersQuery = { __typename?: 'Query', users?: { __typename?: 'UserPaginator', data: Array<{ __typename?: 'User', id: string, name?: string | null, surname?: string | null, email: string, confirmed: boolean, rejected: boolean, userRole: { __typename?: 'UserRole', id: string, roleCode: string, name: string } } | null>, pagination: (
      { __typename?: 'Pagination' }
      & { ' $fragmentRefs'?: { 'PaginationInfoFragment': PaginationInfoFragment } }
    ) } | null };

export const CartItemsFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"CartItems"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"CartResponse"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItemPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"eprelProductCode"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<CartItemsFragment, unknown>;
export const PaginationInfoFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"PaginationInfo"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Pagination"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"totals"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"currentPage"}},{"kind":"Field","name":{"kind":"Name","value":"pageCount"}}]}}]} as unknown as DocumentNode<PaginationInfoFragment, unknown>;
export const UserAddressCollectionFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserAddressCollection"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserAddress"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ID"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"isDefault"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"province"}}]}}]} as unknown as DocumentNode<UserAddressCollectionFragment, unknown>;
export const AddCartDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addCart"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"itemId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"quantity"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"addItemToCart"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"itemId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"itemId"}}},{"kind":"Argument","name":{"kind":"Name","value":"quantity"},"value":{"kind":"Variable","name":{"kind":"Name","value":"quantity"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"CartItems"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"CartItems"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"CartResponse"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItemPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"eprelProductCode"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<AddCartMutation, AddCartMutationVariables>;
export const EditCartDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"editCart"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"cartId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"quantity"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"editCart"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"cartId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"cartId"}}},{"kind":"Argument","name":{"kind":"Name","value":"quantity"},"value":{"kind":"Variable","name":{"kind":"Name","value":"quantity"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"CartItems"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"CartItems"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"CartResponse"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItemPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"eprelProductCode"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<EditCartMutation, EditCartMutationVariables>;
export const EmptyDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"empty"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"emptyCart"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"CartItems"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"CartItems"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"CartResponse"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItemPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"eprelProductCode"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<EmptyMutation, EmptyMutationVariables>;
export const CreateUserBillingDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"CreateUserBilling"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"CreateUserBilling"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUserBilling"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"billingInput"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}}]}}]}}]} as unknown as DocumentNode<CreateUserBillingMutation, CreateUserBillingMutationVariables>;
export const CreateNewOrderDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"createNewOrder"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userAddressId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"newOrder"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}},{"kind":"Argument","name":{"kind":"Name","value":"userAddressId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userAddressId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<CreateNewOrderMutation, CreateNewOrderMutationVariables>;
export const AddAddressDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"addAddress"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UserAddressInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"createUserAddress"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userAddress"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserAddressCollection"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserAddressCollection"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserAddress"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ID"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"isDefault"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"province"}}]}}]} as unknown as DocumentNode<AddAddressMutation, AddAddressMutationVariables>;
export const EditAddressDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"editAddress"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UserAddressInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"editUserAddress"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}},{"kind":"Argument","name":{"kind":"Name","value":"userAddress"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserAddressCollection"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserAddressCollection"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserAddress"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ID"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"isDefault"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"province"}}]}}]} as unknown as DocumentNode<EditAddressMutation, EditAddressMutationVariables>;
export const DeleteAddressDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"deleteAddress"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"deleteUserAddress"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserAddressCollection"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserAddressCollection"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserAddress"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ID"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"isDefault"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"province"}}]}}]} as unknown as DocumentNode<DeleteAddressMutation, DeleteAddressMutationVariables>;
export const ChangeUserStatusDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"changeUserStatus"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userID"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"confirmed"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Boolean"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"rejected"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"Boolean"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"updateUserStatus"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userID"}}},{"kind":"Argument","name":{"kind":"Name","value":"confirmed"},"value":{"kind":"Variable","name":{"kind":"Name","value":"confirmed"}}},{"kind":"Argument","name":{"kind":"Name","value":"rejected"},"value":{"kind":"Variable","name":{"kind":"Name","value":"rejected"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}}]}}]}}]} as unknown as DocumentNode<ChangeUserStatusMutation, ChangeUserStatusMutationVariables>;
export const SearchBrandsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"SearchBrands"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"name"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"searchBrands"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"name"},"value":{"kind":"Variable","name":{"kind":"Name","value":"name"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"quality"}}]}}]}}]} as unknown as DocumentNode<SearchBrandsQuery, SearchBrandsQueryVariables>;
export const UserCartsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"userCarts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"carts"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"totalPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"items"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItemPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<UserCartsQuery, UserCartsQueryVariables>;
export const GetLegalEntityTypesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetLegalEntityTypes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"legalEntityTypes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"isPerson"}}]}}]}}]} as unknown as DocumentNode<GetLegalEntityTypesQuery, GetLegalEntityTypesQueryVariables>;
export const FetchOrderDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"fetchOrder"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"orderId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"order"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"orderId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"priceAmount"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"province"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"userBilling"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"vatNumber"}},{"kind":"Field","name":{"kind":"Name","value":"fiscalCode"}},{"kind":"Field","name":{"kind":"Name","value":"legalEntityType"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orderRows"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}},{"kind":"Field","name":{"kind":"Name","value":"productItemPrice"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItem"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<FetchOrderQuery, FetchOrderQueryVariables>;
export const FetchOrdersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"fetchOrders"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"PaginationInput"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"filter"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"OrderFilterInput"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"ordering"}},"type":{"kind":"ListType","type":{"kind":"NamedType","name":{"kind":"Name","value":"OrderingInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userOrders"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}},{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}}},{"kind":"Argument","name":{"kind":"Name","value":"filter"},"value":{"kind":"Variable","name":{"kind":"Name","value":"filter"}}},{"kind":"Argument","name":{"kind":"Name","value":"ordering"},"value":{"kind":"Variable","name":{"kind":"Name","value":"ordering"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"data"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"status"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"province"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"createdAt"}},{"kind":"Field","name":{"kind":"Name","value":"priceAmount"}},{"kind":"Field","name":{"kind":"Name","value":"orderRows"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"amount"}},{"kind":"Field","name":{"kind":"Name","value":"quantity"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"pagination"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"PaginationInfo"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"PaginationInfo"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Pagination"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"totals"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"currentPage"}},{"kind":"Field","name":{"kind":"Name","value":"pageCount"}}]}}]} as unknown as DocumentNode<FetchOrdersQuery, FetchOrdersQueryVariables>;
export const AllOrdersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"allOrders"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"filter"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"OrdersFilterInput"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"PaginationInput"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"ordering"}},"type":{"kind":"ListType","type":{"kind":"NamedType","name":{"kind":"Name","value":"OrderingInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"allOrders"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"filter"},"value":{"kind":"Variable","name":{"kind":"Name","value":"filter"}}},{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}}},{"kind":"Argument","name":{"kind":"Name","value":"ordering"},"value":{"kind":"Variable","name":{"kind":"Name","value":"ordering"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"data"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"priceAmount"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"userBilling"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"email"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"orderRows"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"amount"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"pagination"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"PaginationInfo"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"PaginationInfo"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Pagination"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"totals"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"currentPage"}},{"kind":"Field","name":{"kind":"Name","value":"pageCount"}}]}}]} as unknown as DocumentNode<AllOrdersQuery, AllOrdersQueryVariables>;
export const GetAllCategoriesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"getAllCategories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productCategories"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"specifications"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"type"}},{"kind":"Field","name":{"kind":"Name","value":"mandatory"}},{"kind":"Field","name":{"kind":"Name","value":"searchable"}}]}}]}}]}}]} as unknown as DocumentNode<GetAllCategoriesQuery, GetAllCategoriesQueryVariables>;
export const SearchDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"search"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"limit"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"offset"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"Int"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"searchInput"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"ProductSearchInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productItems"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"ObjectValue","fields":[{"kind":"ObjectField","name":{"kind":"Name","value":"limit"},"value":{"kind":"Variable","name":{"kind":"Name","value":"limit"}}},{"kind":"ObjectField","name":{"kind":"Name","value":"offset"},"value":{"kind":"Variable","name":{"kind":"Name","value":"offset"}}}]}},{"kind":"Argument","name":{"kind":"Name","value":"productSearchInput"},"value":{"kind":"Variable","name":{"kind":"Name","value":"searchInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"pagination"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"totals"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"currentPage"}},{"kind":"Field","name":{"kind":"Name","value":"pageCount"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productItems"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"supplierQuantity"}},{"kind":"Field","name":{"kind":"Name","value":"price"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}},{"kind":"Field","name":{"kind":"Name","value":"symbol"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"eprelProductCode"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"quality"}}]}},{"kind":"Field","name":{"kind":"Name","value":"vehicleType"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productSpecificationValues"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"specification"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"value"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<SearchQuery, SearchQueryVariables>;
export const ProductItemDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"productItem"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"id"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"productItem"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"id"},"value":{"kind":"Variable","name":{"kind":"Name","value":"id"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"price"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"currency"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"iso_code"}}]}},{"kind":"Field","name":{"kind":"Name","value":"value"}}]}},{"kind":"Field","name":{"kind":"Name","value":"supplierQuantity"}},{"kind":"Field","name":{"kind":"Name","value":"product"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"eprelProductCode"}},{"kind":"Field","name":{"kind":"Name","value":"brand"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"quality"}}]}},{"kind":"Field","name":{"kind":"Name","value":"vehicleType"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"productSpecificationValues"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}},{"kind":"Field","name":{"kind":"Name","value":"specification"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"code"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]}}]} as unknown as DocumentNode<ProductItemQuery, ProductItemQueryVariables>;
export const SearchValuesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"searchValues"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"code"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"value"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"vehicleCode"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"String"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"searchSpecificationValue"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"code"},"value":{"kind":"Variable","name":{"kind":"Name","value":"code"}}},{"kind":"Argument","name":{"kind":"Name","value":"value"},"value":{"kind":"Variable","name":{"kind":"Name","value":"value"}}},{"kind":"Argument","name":{"kind":"Name","value":"vehicleCode"},"value":{"kind":"Variable","name":{"kind":"Name","value":"vehicleCode"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"value"}}]}}]}}]} as unknown as DocumentNode<SearchValuesQuery, SearchValuesQueryVariables>;
export const FetchUserAddressesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"fetchUserAddresses"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userAddress"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"ID"}},{"kind":"Field","name":{"kind":"Name","value":"addressName"}},{"kind":"Field","name":{"kind":"Name","value":"isDefault"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"postalCode"}},{"kind":"Field","name":{"kind":"Name","value":"province"}}]}}]}}]} as unknown as DocumentNode<FetchUserAddressesQuery, FetchUserAddressesQueryVariables>;
export const FetchUserQueryDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"fetchUserQuery"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"userBilling"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"userId"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"province"}},{"kind":"Field","name":{"kind":"Name","value":"cap"}},{"kind":"Field","name":{"kind":"Name","value":"fiscalCode"}},{"kind":"Field","name":{"kind":"Name","value":"vatNumber"}},{"kind":"Field","name":{"kind":"Name","value":"sdiCode"}},{"kind":"Field","name":{"kind":"Name","value":"sdiPec"}},{"kind":"Field","name":{"kind":"Name","value":"legalEntityType"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"email"}}]}}]}}]}}]} as unknown as DocumentNode<FetchUserQueryQuery, FetchUserQueryQueryVariables>;
export const FetchUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"fetchUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"userId"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"ID"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"ID"},"value":{"kind":"Variable","name":{"kind":"Name","value":"userId"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"confirmed"}},{"kind":"Field","name":{"kind":"Name","value":"rejected"}},{"kind":"Field","name":{"kind":"Name","value":"userRole"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"isAdmin"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}},{"kind":"Field","name":{"kind":"Name","value":"userBilling"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine1"}},{"kind":"Field","name":{"kind":"Name","value":"addressLine2"}},{"kind":"Field","name":{"kind":"Name","value":"city"}},{"kind":"Field","name":{"kind":"Name","value":"country"}},{"kind":"Field","name":{"kind":"Name","value":"province"}},{"kind":"Field","name":{"kind":"Name","value":"cap"}},{"kind":"Field","name":{"kind":"Name","value":"fiscalCode"}},{"kind":"Field","name":{"kind":"Name","value":"vatNumber"}},{"kind":"Field","name":{"kind":"Name","value":"sdiCode"}},{"kind":"Field","name":{"kind":"Name","value":"sdiPec"}},{"kind":"Field","name":{"kind":"Name","value":"legalEntityType"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]}}]}}]} as unknown as DocumentNode<FetchUserQuery, FetchUserQueryVariables>;
export const FetchAllUsersDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"fetchAllUsers"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"PaginationInput"}}},{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"filter"}},"type":{"kind":"NamedType","name":{"kind":"Name","value":"UserFilterInput"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"users"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"filter"},"value":{"kind":"Variable","name":{"kind":"Name","value":"filter"}}},{"kind":"Argument","name":{"kind":"Name","value":"pagination"},"value":{"kind":"Variable","name":{"kind":"Name","value":"pagination"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"data"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}},{"kind":"Field","name":{"kind":"Name","value":"surname"}},{"kind":"Field","name":{"kind":"Name","value":"email"}},{"kind":"Field","name":{"kind":"Name","value":"confirmed"}},{"kind":"Field","name":{"kind":"Name","value":"rejected"}},{"kind":"Field","name":{"kind":"Name","value":"userRole"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"roleCode"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}},{"kind":"Field","name":{"kind":"Name","value":"pagination"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"PaginationInfo"}}]}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"PaginationInfo"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"Pagination"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"limit"}},{"kind":"Field","name":{"kind":"Name","value":"totals"}},{"kind":"Field","name":{"kind":"Name","value":"offset"}},{"kind":"Field","name":{"kind":"Name","value":"currentPage"}},{"kind":"Field","name":{"kind":"Name","value":"pageCount"}}]}}]} as unknown as DocumentNode<FetchAllUsersQuery, FetchAllUsersQueryVariables>;