import _ from "lodash";
import config from "../../../../config";
import {ProductCategorySet} from "../../enums/product-specifications-set";

const ProductCategoryNames: { [key in ProductCategorySet]: string } = {
    TYRE: 'tyres'
};

class ProdapiService {
    private baseUrl: string;

    constructor() {
        this.baseUrl = config.prodapi.endpoint;
    }

    private buildUrl( ...path: string[] ): string {
        const string = _.join( path, '/' );

        return `${this.baseUrl}/${string}`;
    }

    private convertCategoryCode(productCategoryCode: ProductCategorySet): string {
        return ProductCategoryNames[productCategoryCode];
    }

    getProductImageUrl( productCode: string, productCategoryCode: ProductCategorySet ): string {
        const categoryName = this.convertCategoryCode( productCategoryCode );

        return this.buildUrl( 'resources', categoryName, productCode );
    }

    getBrandImageUrl( brandCode: string ) {
        brandCode = _.upperCase( _.snakeCase( brandCode ) );

        return this.buildUrl( 'resources', 'brands', brandCode );
    }
}

export default ProdapiService;