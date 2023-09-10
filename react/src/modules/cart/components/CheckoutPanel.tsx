import {faPlus, faTruckFast} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React, {useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {useDispatch, useSelector} from "react-redux";
import Button from "../../../common/components-library/Button";
import {SelectComponent, SelectOption} from "../../../common/components-library/SelectComponent";
import {Currency} from "../../../common/utilities/currency";
import userSelector from "../../user/store/user-selector";
import cartSelector from "../store/cart-selector";
import UserAddressDescriptionList from "./UserAddressDescriptionList";
import {useMutation} from "@apollo/client";
import {NEW_ORDER} from "../../../common/backend/graph/mutation/order";
import {useAuth} from "../../auth/hooks/useAuth";
import {Link, useNavigate} from "react-router-dom";
import {useToast} from "../../../store/toast";
import {ThunkDispatch} from "redux-thunk";
import {emptyCart} from "../store/cart-slice";
import UserAddressModal from "../../user/components/UserAddressModal";
import {useModal} from "../../../common/components/shelly-ui";
import Moment from 'react-moment';
import moment from 'moment-business-days';

const CheckoutPanel: React.FC = () => {
    const userAddresses = useSelector(userSelector.addresses);
    const auth = useAuth();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();
    const [mutate, {loading, error, data}] = useMutation(NEW_ORDER, {
        onCompleted: () => dispatch(emptyCart())
    });
    const {t} = useTranslation();
    const [addressOptions, setOptions] = useState<SelectOption[]>([]);
    const [selectedAddress, setAddress] = useState<SelectOption | null>(null);
    const totalPrice = useSelector(cartSelector.amount);
    const toast = useToast();
    const navigate = useNavigate();

    const getPrice = () => {
        if (!totalPrice || !totalPrice.currency) {
            return "-";
        }

        return Currency.defaultFormat(totalPrice.value, totalPrice.currency?.iso_code as string);
    };

    const getTotalPrice = () => {
        if (!totalPrice || !totalPrice.currency) {
            return "-";
        }

        return Currency.defaultFormat(totalPrice.totalValue, totalPrice.currency?.iso_code as string);
    };


    const getTaxesPrices = () => {
        if (!totalPrice || !totalPrice.currency) {
            return "-";
        }

        return Currency.defaultFormat(totalPrice.taxesValue, totalPrice.currency.iso_code);
    };

    const confirmOrder = () => {
        mutate({
            variables: {
                userId: auth.user?.user?.userID as string,
                userAddressId: selectedAddress?.value.ID
            }
        });
    };

    useEffect(() => {
        if (error) {
            toast.error('Si è verificato un errore' + error.message);
        }
    }, [error]);

    useEffect(() => {
        if (data) {
            toast.info('Ordine creato con successo, procedi con la conferma.');
            navigate('/order/checkout/' + data.newOrder.id);
            return;
        }
    }, [data]);

    useEffect(() => {
        if (!userAddresses || !userAddresses.length) {
            setOptions([]);
            return;
        }

        const options = userAddresses.map((address) => {
            return {
                title: <span> {address.addressName} <span className="text-sm text-accent-content/60"> {address.addressLine1} </span> </span>,
                value: address,
            };
        });


        setOptions(options);
    }, [userAddresses]);

    const modal = useModal();

    return <div className="w-full bg-base-200 rounded-box p-4 flex flex-col">
        <UserAddressModal modal={modal}/>
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
                            placeholder="Seleziona un indirizzo"
                            onChange={setAddress}
                        />
                    </div> : <span> {t("cart.no_address_found")} </span>
            }
        </div>
        <a className="link link-secondary" onClick={modal.open}>
            <FontAwesomeIcon icon={faPlus}/> {t("user_address.add_new_address")}
        </a>
        {selectedAddress &&
            <div className="flex flex-col rounded-box p-4 mt-4 w-full bg-accent-content/10">
                <UserAddressDescriptionList address={selectedAddress.value}/>
            </div>
        }
        <div className="rounded-box p-4 mt-4 w-full bg-accent-content/10">
            <span className="font-semibold"> <FontAwesomeIcon icon={faTruckFast}/> Consegna prevista: </span>
            <Moment className="uppercase" date={moment().businessAdd( 5 )} format="dddd D MMMM"/>
        </div>
        <div className="text-sm mt-4 p-4">
            <div className="flex justify-between"><span>Totale </span> {getPrice()} </div>
            {
                totalPrice && <>
                    <div className="divider"></div>
                    <ul>
                        {totalPrice?.additionsValues?.map((add, key) => <li key={key} className="flex justify-between">
                                <span> {add?.additionName} </span>
                                <span> {Currency.defaultFormat(add?.value as number, totalPrice.currency?.iso_code as string)} </span>
                            </li>
                        )}
                    </ul>
                </>
            }
            <div className="divider"></div>
            <div className="flex justify-between"><span>Totale IVA (22%):</span> {getTaxesPrices()} </div>
        </div>
        <div className="ml-auto mt-10 flex flex-col text-secondary text-sm">
            {t("cart.total_price")}

            <span className="text-4xl font-semibold text-primary">
                {getTotalPrice()}
            </span>
        </div>
        <div className="flex flex-wrap gap-2 ml-auto mt-4">
            <Link
                to="/"
                className="btn btn-outline btn-primary mr-2"
            >
                Continua l'acquisto
            </Link>
            <Button
                type="secondary"
                onClick={confirmOrder}
                loading={loading}
                disabled={!selectedAddress || !totalPrice?.totalValue}
            >
                {t("cart.checkout_button")}
            </Button>
        </div>
    </div>;

};

export default CheckoutPanel;