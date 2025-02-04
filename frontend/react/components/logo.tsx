import Link from 'next/link'
import Image from 'next/image'

export function Logo() {
    return (
        <Link href="/" className="fixed left-4 top-4">
            <Image
                src="/logo.svg"
                alt="SSO Logo"
                width={104}
                height={36}
                className="dark:invert"
                priority
            />
        </Link>
    )
}
