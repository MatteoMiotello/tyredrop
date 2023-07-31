import React from "react";
import {useTranslation} from "react-i18next";
import {UserAddressCollectionFragment} from "../../../__generated__/graphql";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faMapMarkerAlt} from "@fortawesome/free-solid-svg-icons";

type UserAddressDescriptionListProps = {
    address: UserAddressCollectionFragment
}
const UserAddressDescriptionList: React.FC<UserAddressDescriptionListProps> = ( props) => {
    const { t } = useTranslation();

    return <dl>
        <dt className="font-semibold"> <FontAwesomeIcon icon={faMapMarkerAlt}/> Spedisci a: </dt>
        <dd className="ml-2">{props.address.addressName}</dd>
        <dd className="ml-2">{props.address.addressLine1},</dd>
        <dd className="ml-2">{props.address.addressLine2 && props.address.addressLine2}</dd>
        <dd className="ml-2">{props.address.city} ({props.address.province}) {props.address.postalCode} {props.address.country}</dd>
    </dl>;
};

export default UserAddressDescriptionList;