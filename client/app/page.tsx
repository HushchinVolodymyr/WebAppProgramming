import Task1 from "@/components/blocks/task1/task1";
import Task2 from "@/components/blocks/task2/task2";
import Task3 from "@/components/blocks/task3/task3";
import Task4 from "@/components/blocks/task4/task4";
import Task5 from "@/components/blocks/task5/task5";
import Task6 from "@/components/blocks/task6/task6";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs"

export default function Home() {
  return (
    <main className={"flex flex-col items-center h-screen"}>
      <Tabs defaultValue="task6" className="w-screen flex flex-col items-center">
        <TabsList className={`mx-auto`}>
          <TabsTrigger value="task1">Task 1</TabsTrigger>
          <TabsTrigger value="task2">Task 2</TabsTrigger>
          <TabsTrigger value="task3">Task 3</TabsTrigger>
          <TabsTrigger value="task4">Task 4</TabsTrigger>
          <TabsTrigger value="task5">Task 5</TabsTrigger>
          <TabsTrigger value="task6">Task 6</TabsTrigger>
        </TabsList>
        <TabsContent value="task1" className="w-full">
          <Task1 />
        </TabsContent>
        <TabsContent value="task2" className="w-full">
          <Task2 />
        </TabsContent>
        <TabsContent value="task3" className="w-full">
          <Task3 />
        </TabsContent>
        <TabsContent value="task4" className="w-full">
          <Task4 />
        </TabsContent>
        <TabsContent value="task5" className="w-full">
          <Task5 />
        </TabsContent>
        <TabsContent value="task6" className="w-full">
          <Task6 />
        </TabsContent>
      </Tabs>
    </main>
  );
}
