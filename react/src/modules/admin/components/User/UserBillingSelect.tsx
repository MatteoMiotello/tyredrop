import React from "react";
import { UserBillingsQuery, UserBillingsQueryVariables} from "../../../../__generated__/graphql";
import {useQuery} from "../../../../common/backend/graph/hooks";
import {USER_BILLINGS} from "../../../../common/backend/graph/query/users";
import Autocomplete from "../../../../common/components-library/Autocomplete";

type UserBillingSelectProps = {
    name: string
    small?: boolean
}
const UserBillingSelect: React.FC<UserBillingSelectProps> = (props) => {
    const query = useQuery<UserBillingsQuery, UserBillingsQueryVariables>(USER_BILLINGS);

    return <Autocomplete
        getOptions={(search) => {
            if (!search) {
                return undefined;
            }

            return query.refetch({
                name: search
            }).then(res => res.data.userBillings.map((b ) => ({
                        value: b?.id,
                        title: `${b?.name} ${b?.surname} ${b?.vatNumber}`
                    })
                )
            );
        }}
        initialOptions={query.data?.userBillings ? query.data.userBillings.map((b ) => ({
                value: b?.id as string,
                title: `${b?.name} ${b?.surname} [${b?.vatNumber}]`
            })
        ) : []}
        {...props}/>;
};

export default UserBillingSelect;