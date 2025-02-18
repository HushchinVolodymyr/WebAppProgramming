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

const task2FormSchema = z.object({
    coalWeight: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    coalHeatValue: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    oilWeight: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    oilHeatValue: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    gasWeight: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
    gasHeatValue: z.preprocess(toNumber, z.number().min(0, "Must be a number")),
});

export default function Task2() {
    const [task2Result, setTask2Result] = useState<Record<string, any> | null>(null);

    const task2Form = useForm({
        resolver: zodResolver(task2FormSchema),
    });

    function handleTask2Submit(data: z.infer<typeof task2FormSchema>) {
        axios.post('http://localhost:8080/task2', data, {
            headers: { 'Content-Type': 'application/json' },
        })
            .then((response) => {
                console.log('Success:', response.data);
                setTask2Result(response.data);
            })
            .catch((error) => {
                console.error('Error:', error);
            });
    }

    return (
        <section className="flex px-10 gap-6 justify-center">
            <div>
                <Form {...task2Form}>
                    <form onSubmit={task2Form.handleSubmit(handleTask2Submit)} className="flex flex-col items-center w-full">
                        {Object.keys(task2FormSchema.shape).map((field) => (
                            <FormField
                                key={field}
                                control={task2Form.control}
                                name={field as keyof typeof task2FormSchema.shape}
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
                            Calculate Task 2
                        </Button>
                    </form>
                </Form>
            </div>

            <div>
                <h1 className="text-xl">Results:</h1>
                {task2Result ? (
                    <div className="border p-4 rounded">
                        {Object.entries(task2Result).map(([key, value]) => (
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
