import {faCheck} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Listbox, Transition} from "@headlessui/react";
import React, {Fragment, ReactNode, useEffect, useState} from "react";
import {PropsWithValidators} from "../validation/validators";

export type SelectOption = {
    title: string | ReactNode
    value: any
    disabled?: boolean
}

interface SelectProps extends PropsWithValidators<SelectOption | null> {
    className?: string
    options: (SelectOption | null)[]
    placeholder?: string
    name: string
    onChange?: (value: any) => void
    defaultValue?: SelectOption
}

export const SelectComponent: React.FC<SelectProps> = (props: SelectProps) => {
    const [selected, setSelected] = useState(props.defaultValue ?? props.options[0] ?? null);
    const [ error, setError ] = useState<string | null>( null );

    useEffect( () => {
        if ( props.options ) {
            setSelected( props.options[0] );
        }
    }, [props.options] );

    useEffect( () => {
        if ( !props.validators ) {
            return;
        }

        props.validators.every( ( validator ) => {
            const error = validator( selected );

            if ( error ) {
                setError( error );
                return false;
            }
            return true;
        } );

    }, [selected] );

    useEffect( () => {
        if ( props.onChange ) {
            props.onChange(selected);
        }
    }, [selected] );

    return <Listbox value={selected?.value || null} onChange={setSelected} name={props.name}>
        <div className={"relative " + props.className}>
            <Listbox.Button className="select relative w-full cursor-default flex items-center">
                <span className="truncate">{ selected ? selected.title : ( props.placeholder ?? '' ) }</span>
            </Listbox.Button>
            { error ? <span className="label-text-alt text-error">{error}</span> : ''}
            <Transition
                as={Fragment}
                leave="transition ease-in duration-100"
                leaveFrom="opacity-100"
                leaveTo="opacity-0"
            >
                <Listbox.Options className="absolute mt-1 pl-0 max-h-60 w-full overflow-auto rounded-md bg-white py-1 text-base shadow-lg ring-1 ring-black ring-opacity-5 focus:outline-none sm:text-sm">
                    {props.options.map((option, key) => (
                        <Listbox.Option
                            key={key}
                            className={({ active }) =>
                                `relative cursor-default select-none py-2 list-none ${
                                    active ? 'bg-base-200 text-secondary' : 'text-gray-900'
                                }`
                            }
                            value={option}
                        >
                            {({ selected }) => (
                                <>
                      <span
                          className={`block truncate ml-8 ${
                              selected ? 'font-medium' : 'font-normal'
                          }`}
                      >
                        {option ? option.title : ''}
                      </span>
                                    {selected ? (
                                        <span className="absolute inset-y-0 left-0 flex items-center text-primary ml-2">
                          <FontAwesomeIcon icon={faCheck} className="h-5 w-5" aria-hidden="true" />
                        </span>
                                    ) : null}
                                </>
                            )}
                        </Listbox.Option>
                    ))}
                </Listbox.Options>
            </Transition>
        </div>
    </Listbox>;
};