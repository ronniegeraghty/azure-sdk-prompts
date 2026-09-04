interface HyokaLogoProps {
  className?: string;
  size?: number;
}

export function HyokaLogo({ className = "", size = 24 }: HyokaLogoProps) {
  return (
    <svg
      width={size}
      height={size}
      viewBox="0 0 32 32"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
      className={className}
      role="img"
      aria-label="hyoka logo"
      focusable="false"
    >
      {/* Stylized checkmark forming "h" */}
      <path
        d="M8 16L12 20L24 8"
        stroke="currentColor"
        strokeWidth="3"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      {/* Rising bar chart elements */}
      <rect x="5" y="22" width="3" height="5" rx="1" fill="currentColor" opacity="0.6" />
      <rect x="10" y="19" width="3" height="8" rx="1" fill="currentColor" opacity="0.7" />
      <rect x="15" y="16" width="3" height="11" rx="1" fill="currentColor" opacity="0.8" />
      <rect x="20" y="13" width="3" height="14" rx="1" fill="currentColor" opacity="0.9" />
      <rect x="25" y="10" width="3" height="17" rx="1" fill="currentColor" opacity="1" />
    </svg>
  );
}
