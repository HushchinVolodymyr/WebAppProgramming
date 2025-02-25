"use client";

import { useState } from "react";
import { z } from "zod";
import {
    Form,
    FormControl,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import axios from "axios";

const toNumber = (val: unknown) =>
    typeof val === "string" && val.trim() !== "" ? parseFloat(val) : undefined;

const task4FormSchema = z.object({
    transformerPowerCable: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    voltageCable: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    distanceCable: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    maxVoltDropCable: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    circuitVoltage: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    circuitImpedance: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    voltageSubstation: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    rNormal: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    xNormal: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    rMin: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    xMin: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
});

export default function Task4() {
    const [task4Result, setTask4Result] = useState<Record<string, any> | null>(null);

    const task4Form = useForm({
        resolver: zodResolver(task4FormSchema),
    });

    function handleTask4Submit(data: z.infer<typeof task4FormSchema>) {
        axios.post('http://localhost:8080/task4', data, {
            headers: { 'Content-Type': 'application/json' },
        })
            .then((response) => {
                console.log('Success:', response.data);
                setTask4Result(response.data);
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    return (
        <section className="flex px-10 gap-6 justify-center">
            <div>
                <Form {...task4Form}>
                    <form onSubmit={task4Form.handleSubmit(handleTask4Submit)} className="flex flex-col items-center w-full">
                        {Object.keys(task4FormSchema.shape).map((field) => (
                            <FormField
                                key={field}
                                control={task4Form.control}
                                name={field as keyof typeof task4FormSchema.shape}
                                render={({ field }) => (
                                    <FormItem>
                                        <FormLabel>{field.name.charAt(0).toUpperCase() + field.name.slice(1)}</FormLabel>
                                        <FormControl>
                                            <Input
                                                type="number"
                                                step="any"
                                                placeholder={field.name}
                                                {...field}
                                                value={field.value ?? ""}
                                                onChange={(e) => {
                                                    const val = e.target.value;
                                                    field.onChange(val === "" ? undefined : val);
                                                }}
                                                style={{ width: "300px" }}
                                            />
                                        </FormControl>
                                        <FormMessage />
                                    </FormItem>
                                )}
                            />
                        ))}

                        <Button type="submit" className="mt-2">
                            Calculate Task 3
                        </Button>
                    </form>
                </Form>
            </div>

            <div className="w-full max-w-3xl">
                <h1 className="text-xl font-bold mb-4 text-center">Results:</h1>
                {task4Result ? (
                    <div className="border rounded-lg overflow-hidden shadow-lg">
                        <table className="w-full border-collapse">
                            <thead>
                                <tr>
                                    <th className="border p-2 text-left">Param</th>
                                    <th className="border p-2 text-left">Value</th>
                                </tr>
                            </thead>
                            <tbody>
                                {Object.entries(task4Result).map(([key, value], index) => (
                                    <tr key={key} >
                                        <td className="border p-2">{key}</td>
                                        <td className="border p-2">
                                            {key === "CableOptions" ? (
                                                <div className="flex flex-col gap-1">
                                                    {Array.isArray(value) && value.map((cable, index) => (
                                                        <div key={index} className="grid grid-cols-5 gap-2 p-2 border-b">
                                                            <span className="font-medium">{cable.Section} mm²</span>
                                                            <span>{cable.Material}</span>
                                                            <span>{cable.MaxCurrent} A</span>
                                                            <span>{cable.Resistance} Ω</span>
                                                            <span>{cable.Reactance} Ω</span>
                                                        </div>
                                                    ))}
                                                </div>
                                            ) : null}
                                            {key === "SubstationStates" ? (
                                                <div className="flex flex-col gap-1">
                                                    {Object.entries(value).map(([state, val], index) => (
                                                        <div key={index} className="grid grid-cols-2 gap-2 p-2 border-b">
                                                            <span className="font-medium">{state}</span>
                                                            <span>{val}</span>
                                                        </div>
                                                    ))}
                                                </div>
                                            ) : null}
                                            {key !== "CableOptions" && key !== "SubstationStates" ? (
                                                <span>{value}</span>
                                            ) : null}
                                        </td>

                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    </div>
                ) : (
                    <p className="text-gray-500 text-center">No Data</p>
                )}
            </div>
        </section>
    );
}
