import React, { useState, useEffect } from "react";
import type { Area } from "../api/report";
import { getAreas} from "../api/report";

interface FilterFormProps {
  onSubmit: (area_id?: string, dateFrom?: string, dateTo?: string) => void;
}

export default function FilterForm({ onSubmit }: FilterFormProps) {
  const [areas, setAreas] = useState<Area[]>([]);
  const [loading, setLoading] = useState(false);

  // Fetch area saat komponen load
  useEffect(() => {
    async function fetchAreas() {
      setLoading(true);
      try {
        const data = await getAreas();
        setAreas(data);
      } catch (err) {
        console.error("Failed to fetch areas", err);
      } finally {
        setLoading(false);
      }
    }
    fetchAreas();
  }, []);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        const target = e.target as typeof e.target & {
          area_id: { value: string };
          dateFrom: { value: string };
          dateTo: { value: string };
        };
        onSubmit(target.area_id.value, target.dateFrom.value, target.dateTo.value);
      }}
      style={{
        display: "flex",
        gap: "15px",
        justifyContent: "center",
        alignItems: "center",
        flexWrap: "wrap",
      }}
    >
      <label>
        Select Area:
        <select name="area_id" style={{ marginLeft: 5 }}>
          <option value="">-- All Areas --</option>
          {loading ? (
            <option disabled>Loading...</option>
          ) : (
            areas.map((area) => (
              <option key={area.id} value={area.id}>
                {area.name}
              </option>
            ))
          )}
        </select>
      </label>

      <label>
        Date From:
        <input type="date" name="dateFrom" style={{ marginLeft: 5 }} />
      </label>

      <label>
        Date To:
        <input type="date" name="dateTo" style={{ marginLeft: 5 }} />
      </label>

      <button
        type="submit"
        style={{
          padding: "8px 18px",
          fontSize: "14px",
          cursor: "pointer",
        }}
      >
        View
      </button>
    </form>
  );
}
