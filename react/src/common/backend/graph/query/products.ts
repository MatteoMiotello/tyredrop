import {gql} from "../../../../__generated__";

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

export const SEARCH_PRODUCTS = gql(`
    query search($limit: Int!, $offset: Int!, $searchInput: ProductSearchInput) {
     productItems(
          pagination: { limit: $limit, offset: $offset }
          productSearchInput: $searchInput
     ) {
          pagination {
               limit
               totals
               offset
               currentPage
               pageCount
          }
          productItems {
               id
               supplierQuantity
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
                   vehicleType {
                       name
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
`);

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
        supplierQuantity
        product {
            code
            name
            eprelProductCode
            brand {
                name
                code
                quality
            }
            vehicleType {
                name
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

export const SEARCH_PRODUCT_SPECIFICATION_VALUES = gql( /*GraphQL*/`
    query searchValues( $code: String!, $value: String, $vehicleCode: String) {
        searchSpecificationValue( code: $code, value: $value, vehicleCode: $vehicleCode ) {
            id
            specification {
                code 
            }
            value
        }
    }
`);

export const PRICE_MARKUPS = gql( /* GraphQL */ `
    query priceMarkups {
        priceMarkups {
            id
            markupPercentage
            productCategory {
                id
                name
            }
            product {
                id
                name
            }
            brand {
                id
                code
                name
            }
            productSpecificationValue {
                id
                value
                specification {
                    name
                }
            }
        }
    }
` );