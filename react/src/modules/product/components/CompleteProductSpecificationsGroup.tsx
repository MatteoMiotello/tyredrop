import {ColumnDef} from "@tanstack/react-table";
import React from "react";
import {ProductSpecificationValue} from "../../../__generated__/graphql";
import {BasicTable, useTable} from "../../../common/components/shelly-ui";
import {specificationProvider} from "../specifications/provider";

type CompleteProductSpecificationsGroupProps = {
    specifications: ProductSpecificationValue[]
}

const CompleteProductSpecificationsGroup: React.FC<CompleteProductSpecificationsGroupProps> = ( props ) => {
    const columns: ColumnDef<ProductSpecificationValue>[] = [
        {
            accessorKey: "specification",
            header: "Specifica",
            cell: ( props ) => <span className="text-sm text-secondary font-semibold flex-1"> { props.row.original.specification.name } </span>
        },
        {
            accessorKey: "value",
            header: "Valore",
            cell: ( props ) => specificationProvider( props.row.original as ProductSpecificationValue).content
        }
    ];

    const {table} = useTable({
        data: props.specifications,
        columns: columns,
    });
    
    return <BasicTable table={table}/>;
};

export default CompleteProductSpecificationsGroup;