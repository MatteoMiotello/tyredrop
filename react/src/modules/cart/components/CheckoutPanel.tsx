import {faPlus} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {useSelector} from "react-redux";
import Button from "../../../common/components-library/Button";
import Modal from "../../../common/components-library/Modal";
import {SelectComponent, SelectOption} from "../../../common/components-library/SelectComponent";
import {Currency} from "../../../common/utilities/currency";
import useModal from "../../../hooks/useModal";
import UserAddressForm from "../../user/components/UserAddressForm";
import userSelector from "../../user/store/user-selector";
import cartSelector from "../store/cart-selector";
import UserAddressDescriptionList from "./UserAddressDescriptionList";

const CheckoutPanel: React.FC = () => {
    const userAddresses = useSelector(userSelector.addresses);
    const {t} = useTranslation();
    const [addressOptions, setOptions] = useState<SelectOption[]>([]);
    const [selectedAddress, setAddress] = useState<SelectOption | null>(null);
    const totalPrice = useSelector(cartSelector.amount);

    const getPrice = () => {
        if (!totalPrice || !totalPrice.currency) {
            return "-";
        }

        return Currency.defaultFormat(totalPrice.value, totalPrice.currency?.iso_code as string);
    };

    useEffect(() => {
        if (!userAddresses || !userAddresses.length) {
            setOptions([
                {
                    title: t("user_address.no_address_found"),
                    value: 1,
                    disabled: true
                }
            ]);
            return;
        }

        const options = userAddresses.map((address) => {
            return {
                title: <span> {address.addressName} <span className="text-sm text-accent-content/60 "> {address.addressLine1} </span> </span>,
                value: address,
            };
        });


        setOptions(options);
    }, [userAddresses]);

    const {openModal, closeModal} = useModal(<Modal.Content>
        <Modal.Header>
            {t("user_address.new_address_modal_title")}
        </Modal.Header>
        <UserAddressForm onSuccess={() => closeModal()}>
            <Modal.Action>
                <Button onClick={() => closeModal()} htmlType="button">
                    {t("user_address.close_modal")}
                </Button>
                <Button type="primary" htmlType="submit">
                    {t("user_address.submit_form")}
                </Button>
            </Modal.Action>
        </UserAddressForm>
    </Modal.Content>);

    return <div className="w-full bg-base-200 rounded-box p-4 flex flex-col">
        <div className="mb-4">
            {
                (userAddresses && userAddresses.length) ?
                    <div>

                        <label className="label label-text">
                            {t("cart.select_address_field")}
                        </label>
                        <SelectComponent
                            options={addressOptions}
                            name="user_address"
                            onChange={setAddress}
                        />
                    </div> : <span> {t("cart.no_address_found")} </span>
            }
        </div>
        <a className="link link-secondary" onClick={() => openModal()}>
            <FontAwesomeIcon icon={faPlus}/> {t("user_address.add_new_address")}
        </a>
        {selectedAddress &&
            <div className="flex flex-col rounded-box p-4 mt-4 w-full bg-accent-content/10">
                <UserAddressDescriptionList address={selectedAddress.value}/>
            </div>
        }
        <div className="ml-auto mt-10 flex flex-col text-secondary text-sm">
            {t("cart.total_price")}

            <span className="text-4xl font-semibold text-primary">
                {getPrice()}
            </span>
        </div>
        <Button className="ml-auto mt-4" type="secondary" onClick={() => openModal()}>
            {t("cart.checkout_button")}
        </Button>
    </div>;

};

export default CheckoutPanel;