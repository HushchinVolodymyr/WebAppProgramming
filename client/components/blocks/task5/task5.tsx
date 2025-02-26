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

const task5FormSchema = z.object({
    egActivity: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    pl: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    plLong: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    transmission: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    activator: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    connection: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    connectionTimes: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    costAvaria: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    costPlan: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
});

export default function Task5() {
    const [task5Result, setTask5Result] = useState<Record<string, any> | null>(null);

    const task5Form = useForm({
        resolver: zodResolver(task5FormSchema),
    });

    function handleTask5Submit(data: z.infer<typeof task5FormSchema>) {
        axios.post('http://localhost:8080/task5', data, {
            headers: { 'Content-Type': 'application/json' },
        })
            .then((response) => {
                console.log('Success:', response.data);
                setTask5Result(response.data);
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    return (
        <section className="flex px-10 gap-6 justify-center">
            <div>
                <Form {...task5Form}>
                    <form onSubmit={task5Form.handleSubmit(handleTask5Submit)} className="flex flex-col items-center w-full">
                        {Object.keys(task5FormSchema.shape).map((field) => (
                            <FormField
                                key={field}
                                control={task5Form.control}
                                name={field as keyof typeof task5FormSchema.shape}
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
                {task5Result ? (
                    <div className="border rounded-lg overflow-hidden shadow-lg">
                        <table className="w-full border-collapse">
                            <thead>
                                <tr>
                                    <th className="border p-2 text-left">Param</th>
                                    <th className="border p-2 text-left">Value</th>
                                </tr>
                            </thead>
                            <tbody>
                                {Object.entries(task5Result).map(([key, value], index) => (
                                    <tr key={key} >
                                        <td className="border p-2">{key}</td>
                                        <td className="border p-2">{value}</td>
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
