import React from "react";
import type { TableItem } from "../api/report";

interface DataTableProps {
  data: TableItem[];
}

export default function DataTable({ data }: DataTableProps) {
  const areas = Array.from(new Set(data.map((d) => d.area)));

  const brands = Array.from(new Set(data.map((d) => d.brand)));

  const lookup: Record<string, Record<string, number>> = {};

  data.forEach(({ brand, area, value }) => {
    if (!lookup[brand]) lookup[brand] = {};
    lookup[brand][area] = value;
  });

  return (
    <table border={1} cellPadding={8} style={{ width: "100%", borderCollapse: "collapse" }}>
      <thead>
        <tr>
          <th>Brand</th>
          {areas.map((area) => (
            <th key={area}>{area}</th>
          ))}
        </tr>
      </thead>
      <tbody>
        {brands.map((brand) => (
          <tr key={brand}>
            <td>{brand}</td>
            {areas.map((area) => (
              <td key={area}>{lookup[brand]?.[area]?.toFixed(0) ?? "-"}</td>
            ))}
          </tr>
        ))}
      </tbody>
    </table>
  );
}
