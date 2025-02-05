import Link from 'next/link'
import Image from 'next/image'

export function Logo() {
    return (
        <Link href="/" className="fixed left-4 top-4">
            <Image
                src="/Logo.svg"
                alt="SSO Logo"
                width={139}
                height={48}
                className="dark:invert"
                priority
            />
        </Link>
    )
}
