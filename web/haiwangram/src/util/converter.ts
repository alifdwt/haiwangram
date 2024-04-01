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

export { TimeConverter };
