import {gql} from "@apollo/client";

export const ALL_CATEGORIES_WITH_SPECIFICATIONS = gql(`
    query getAllCategories {
        productCategories {
            id
            name
            code
            specifications {
                id
                code
                name
                type
                mandatory
                searchable
            }
        }
    }
`);

export const SEARCH_PRODUCTS = gql( `
    query search($limit: Int!, $offset: Int!, $searchInput: ProductSearchInput) {
     productItems(
          pagination: { limit: $limit, offset: $offset }
          productSearchInput: $searchInput
     ) {
          pagination {
               limit
               totals
               offset
          }
          productItems {
               id
               price {
                    value
                    currency {
                         iso_code
                         symbol
                    }
               }
               product {
                    id
                    name
                    code
                    eprelProductCode
                    brand {
                         name
                         code
                         quality
                    }
                    productSpecificationValues {
                        specification {
                            code
                        }
                        value
                    }
               }
          }
     }
}
` );

export const PRODUCT_ITEM = gql(`
    query productItem( $id: ID! ) {
    productItem( id: $id ) {
        id
        price {
            currency {
                iso_code
            }
            value
        }
        product {
            code
            name
            eprelProductCode
            brand {
                name
                code
                quality
            }
            productSpecificationValues {
                value
                specification {
                    code
                    name
                }
            }
        }
    }
}
`);