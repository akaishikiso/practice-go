"use client";

import { useEffect, useState } from "react";

type Menu = {
  id: number;
  name: string;
  date: string;
  note: string;
};

export default function Menu() {
  const [menus, setMenus] = useState<Menu[]>([]);
  const [name, setName] = useState("");
  const [date, setDate] = useState("");
  const [note, setNote] = useState("");

  useEffect(() => {
    fetchMenus();
  }, []);

  async function fetchMenus() {
    const res = await fetch("http://localhost:8080/menus");
    const data = await res.json();
    setMenus(data);
  }

  async function addMenu() {
    await fetch("http://localhost:8080/menus", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ name, date, note }),
    });
    setName("");
    setDate("");
    setNote("");
    fetchMenus();
  }

  return (
    <main className="p-4">
      <h1 className="text-xl font-bold mb-2">献立一覧</h1>
      <ul className="mb-4">
        {menus.map((menu) => (
          <li key={menu.id}>
            {menu.date}: {menu.name} - {menu.note}
          </li>
        ))}
      </ul>

      <div className="mb-2">
        <input
          value={name}
          onChange={(e) => setName(e.target.value)}
          placeholder="献立名"
        />
        <input
          value={date}
          onChange={(e) => setDate(e.target.value)}
          placeholder="日付"
        />
        <input
          value={note}
          onChange={(e) => setNote(e.target.value)}
          placeholder="メモ"
        />
        <button onClick={addMenu}>追加</button>
      </div>
    </main>
  );
}
