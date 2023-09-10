import {faCheck, faChevronDown} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Combobox, Transition} from "@headlessui/react";
import React, {ChangeEventHandler, Fragment, ReactNode, useEffect, useState} from "react";
import {useTranslation} from "react-i18next";
import {PropsWithValidators, ValidationHandler} from "../validation/validators";
import _ from "lodash";

export type AutocompleteQueryHandler = (query: string) => Promise<AutocompleteOption[] | null> | undefined

interface AutocompleteProps<T = any> extends PropsWithValidators {
    getOptions: AutocompleteQueryHandler | undefined;
    initialOptions: AutocompleteOption<T>[];
    className?: string;
    name: string;
    placeholder?: string | undefined;
    labelText?: string | undefined | null;
    defaultValue?: T
    disable?: boolean
    small?: boolean
}

export type AutocompleteOption<T = any> = {
    title: ReactNode,
    content?: ReactNode
    value: T
}

const Autocomplete: React.FC<AutocompleteProps> = (props) => {
    const [selectedValue, setSelectedValue] = useState<any>(props.defaultValue || "");
    const [query, setQuery] = useState('');
    const [filteredOptions, setFilteredOptions] = useState(props.initialOptions);
    const [error, setError] = useState<string | null>(null);
    const {t} = useTranslation();

    useEffect(() => {
        if (props.defaultValue) {
            setSelectedValue(props.defaultValue);
        }
    }, [props.defaultValue]);

    const performQuery = () => {
        if ( !props.getOptions ) {
            return;
        }

        const options = props.getOptions(query);

        if ( !options ) {
            return;
        }

        options.then(res => {
            let options: AutocompleteOption[] = [];
            if (res) {
                options = res;
            }

            if (query.length && !options.length) {
                options = [
                    {
                        title: query,
                        value: query
                    },
                    ...options
                ];
            }
            setFilteredOptions(_.uniqWith(options, _.isEqual));
        })
            .catch((err) => {
                return null;
            });
    };

    useEffect(() => {
        performQuery();
    }, [query]);

    const findOption = (value: any) => {
        if (!filteredOptions.length) {
            return null;
        }

        if (!value) {
            return null;
        }

        return filteredOptions.find(opt => opt.value == value);
    };

    const onChange: ChangeEventHandler<HTMLInputElement> = (event) => {
        const value = event.target.value;

        if (!value) {
            setSelectedValue(null);
        }

        if (!props.validators) {
            return;
        }

        props.validators?.every((validator: ValidationHandler) => {
            const error = validator(value as string);
            setError(error);

            return !error;
        });
    };

    return <Combobox value={selectedValue} onChange={setSelectedValue} name={props.name} disabled={props.disable}>
        <div className={"relative " + props.className}>
            {props.labelText &&
                <label className="label">
                        <span className="label-text">
                            {props.labelText}
                        </span>
                </label>
            }
            <div
                className={`select ${props.small && "select-sm text-sm"} select-bordered relative w-full cursor-default overflow-hidden p-0 ${error ? 'select-error' : ''}`}>
                <Combobox.Input
                    autoComplete="off"
                    className={`input w-full border-none h-full font-normal flex items-center ${props.small && 'text-sm'} `}
                    displayValue={(value: any) => {
                        const option = findOption(value);
                        return option ? option.title : value;
                    }}
                    onChange={(event) => {
                        setQuery(event.target.value);
                        onChange(event);
                    }}
                    onFocus={onChange}
                    placeholder={props.placeholder}
                />
                <Combobox.Button className="absolute inset-y-0 right-0 flex items-center pr-2 border-none"
                                 onClick={performQuery}>
                    {
                        filteredOptions.length ?
                        <FontAwesomeIcon icon={faChevronDown}
                                         className="h-5 w-5 text-primary label-text"
                                         aria-hidden="true"
                        /> : null
                    }
                </Combobox.Button>
            </div>
            {error ? <span className="label-text-alt text-error">{error}</span> : ''}
            <Transition
                as={Fragment}
                leave="transition ease-in duration-100"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
                afterLeave={() => setQuery('')}
            >
                <Combobox.Options
                    className="absolute z-30 mt-1 pl-0 max-h-96 w-full overflow-auto rounded-md bg-base-100 py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    {
                        filteredOptions.map((option, key) => (
                            <Combobox.Option
                                key={key}
                                className={({active}) =>
                                    `relative cursor-default select-none py-2 list-none ${
                                        active ? 'bg-base-200 text-secondary' : 'text-gray-900'
                                    }`
                                }
                                value={option.value}
                            >
                                {({selected, active}) => (
                                    <>
                        <span
                            className={`block truncate ml-8 ${props.small && 'text-sm'} ${
                                selected ? 'font-medium' : 'font-normal'
                            }`}
                        >
                          {option.content ? option.content : option.title}
                        </span>
                                        {selected ? (
                                            <span
                                                className={`absolute ${props.small && 'text-sm'} inset-y-0 left-0 flex items-center text-primary ml-2 ${
                                                    active ? 'text-primary' : 'text-grey-600'
                                                }`}
                                            >
                            <FontAwesomeIcon icon={faCheck} className="h-5 w-5" aria-hidden="true"/>
                          </span>
                                        ) : null}
                                    </>
                                )}
                            </Combobox.Option>
                        ))
                    }
                </Combobox.Options>
            </Transition>
        </div>
    </Combobox>;
};

export default Autocomplete;