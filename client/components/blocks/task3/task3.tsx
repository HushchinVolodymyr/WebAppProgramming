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

const task3FormSchema = z.object({
    energyPower: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    oldErrorDeviation: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    newErrorDeviation: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    energyPrice: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
});

export default function Task3() {
    const [task3Result, setTask3Result] = useState<Record<string, any> | null>(null);

    const task3Form = useForm({
        resolver: zodResolver(task3FormSchema),
    });

    function handleTask3Submit(data: z.infer<typeof task3FormSchema>) {
        axios.post('http://localhost:8080/task3', data, {
            headers: { 'Content-Type': 'application/json' },
        })
            .then((response) => {
                console.log('Success:', response.data);
                setTask3Result(response.data);
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    return (
        <section className="flex px-10 gap-6 justify-center">
            <div>
                <Form {...task3Form}>
                    <form onSubmit={task3Form.handleSubmit(handleTask3Submit)} className="flex flex-col items-center w-full">
                        {Object.keys(task3FormSchema.shape).map((field) => (
                            <FormField
                                key={field}
                                control={task3Form.control}
                                name={field as keyof typeof task3FormSchema.shape}
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

            <div>
                <h1 className="text-xl">Results:</h1>
                {task3Result ? (
                    <div className="border p-4 rounded">
                        {Object.entries(task3Result).map(([key, value]) => (
                            <div key={key} className="flex justify-between py-1">
                                <span className="font-medium">{key}:</span>
                                <span>{value}</span>
                            </div>
                        ))}
                    </div>
                ) : null}
            </div>
        </section>
    );
}
