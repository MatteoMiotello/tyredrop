import {faCheck} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {Listbox, Transition} from "@headlessui/react";
import React, {Fragment, useState} from "react";

type SelectOption = {
    title: string
    value: any
    disabled?: boolean
}

type SelectProps = {
    className?: string
    options: SelectOption[]
}

export const SelectComponent: React.FC<SelectProps> = (props: SelectProps) => {
    if (!props.options || props.options.length == 0) {
        throw new Error('No options prompted');
    }

    const [selected, setSelected] = useState(props.options[0]);

    return <Listbox value={selected} onChange={setSelected}>
        <div className={"relative mt-1 " + props.className}>
            <Listbox.Button className="select relative w-full cursor-default flex items-center">
                <span className="truncate">{selected.title}</span>
            </Listbox.Button>
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
                        {option.title}
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