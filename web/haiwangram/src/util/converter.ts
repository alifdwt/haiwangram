import {
  differenceInMinutes,
  differenceInHours,
  differenceInDays,
  differenceInWeeks,
  differenceInMonths,
  differenceInYears,
  parseISO,
} from "date-fns";

const TimeConverter = (time: string) => {
  // 2024-04-01T00:06:21.424178+07:00 => 01 Apr 00:06
  const date = new Date(time);
  const options: Intl.DateTimeFormatOptions = {
    year: "numeric",
    month: "short",
    day: "numeric",
    hour: "numeric",
    minute: "numeric",
  };

  return date.toLocaleString("id-ID", options);
};

function getIndonesianDuration(timestamp: string): string {
  const currentDate = new Date();
  const targetDate = parseISO(timestamp);

  const minutesDifference = differenceInMinutes(currentDate, targetDate);
  const hoursDifference = differenceInHours(currentDate, targetDate);
  const daysDifference = differenceInDays(currentDate, targetDate);
  const weeksDifference = differenceInWeeks(currentDate, targetDate);
  const monthsDifference = differenceInMonths(currentDate, targetDate);
  const yearsDifference = differenceInYears(currentDate, targetDate);

  if (minutesDifference < 60) {
    return `${minutesDifference} menit lalu`;
  } else if (hoursDifference < 24) {
    return `${hoursDifference} jam lalu`;
  } else if (daysDifference === 1) {
    return `kemarin`;
  } else if (daysDifference < 7) {
    return `${daysDifference} hari lalu`;
  } else if (weeksDifference === 1) {
    return `1 minggu lalu`;
  } else if (weeksDifference < 4) {
    return `${weeksDifference} minggu lalu`;
  } else if (monthsDifference === 1) {
    return `1 bulan lalu`;
  } else if (monthsDifference < 12) {
    return `${monthsDifference} bulan lalu`;
  } else if (yearsDifference === 1) {
    return `1 tahun lalu`;
  } else {
    return `${yearsDifference} tahun lalu`;
  }
}

export { TimeConverter, getIndonesianDuration };
